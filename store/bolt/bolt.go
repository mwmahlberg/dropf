package bolt

import (
	"bytes"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jubalh/dropf/store"
	"github.com/pborman/uuid"
)

var (
	defaultFileBucket   = []byte("files")
	defaultNameBucket   = []byte("names")
	defaultStatusBucket = []byte("isPublic")
	defaultExpireBucket = []byte("expires")
	defaultFileName     = "dropf.db"
)

// Bolt implements the FileStorer interface.
//
// As the name implies, it uses BoltDB as the underlying storage.
//
//    https://github.com/boltdb/bolt
type Bolt struct {
	db *bolt.DB
	fb []byte
	nb []byte
	sb []byte
	eb []byte
}

func (s *Bolt) init() error {
	return s.db.Update(
		func(tx *bolt.Tx) error {
			var err error

			for _, bucketname := range [][]byte{s.eb, s.fb, s.nb, s.sb} {
				if _, err = tx.CreateBucketIfNotExists(bucketname); err != nil {
					return err
				}
			}

			return nil
		})
}

// FileBucket sets the name of the BoltDB bucket the files are stored in
func FileBucket(name string) func(b *Bolt) {
	return func(b *Bolt) {
		b.fb = []byte(name)
	}
}

// FilenameBucket sets the name of the BoltDB bucket the filenames are stored in.
func FilenameBucket(name string) func(b *Bolt) {
	return func(b *Bolt) {
		b.nb = []byte(name)
	}
}

// StatusBucket sets the name of the BoltDB bucket storing wether a file is public or not.
func StatusBucket(name string) func(b *Bolt) {
	return func(b *Bolt) {
		b.sb = []byte(name)
	}
}

// NewBoltFileStore creates a new FileStore.
// If db is nil, a new database file named dropf.db is created in the current directory.
// You may pass any number of the bolt.*Options to configure the store.
//
// Returns either a Bolt store and a nil error or an nil store and the error which occured
// during setup.
func NewBoltFileStore(db *bolt.DB, options ...func(b *Bolt)) (b *Bolt, err error) {

	if db == nil {
		if db, err = bolt.Open(defaultFileName, 0600, nil); err != nil {
			return nil, err
		}

	}

	b = &Bolt{
		db: db,
		eb: defaultExpireBucket,
		fb: defaultFileBucket,
		nb: defaultNameBucket,
		sb: defaultStatusBucket,
	}

	for _, option := range options {
		option(b)
	}

	if err := b.init(); err != nil {
		panic(err)
	}
	return b, nil
}

// Store implements the Store method of the FileStorer interface.
//
// It takes the file and stores it's content and metadata in a BoltDB data file.
func (s Bolt) Store(file *store.File) (id uuid.UUID, err error) {
	if file.Metadata.ID == nil {
		file.Metadata.ID = uuid.NewRandom()
	}

	err = s.db.Update(
		func(tx *bolt.Tx) error {
			buf := new(bytes.Buffer)

			buf.ReadFrom(file.Content)

			if err = tx.Bucket(s.fb).Put(file.Metadata.ID, buf.Bytes()); err != nil {
				return err
			}

			if err = tx.Bucket(s.nb).Put(file.Metadata.ID, []byte(file.Metadata.Name)); err != nil {
				return err
			}

			if file.Metadata.Public {
				if err = tx.Bucket(s.sb).Put(file.Metadata.ID, []byte("")); err != nil {
					return err
				}
			}

			if file.Metadata.Expires.After(time.Time{}) {
				if file.Metadata.Expires.Before(time.Now()) {
					return store.ExpirationInPastError(file.Metadata.Expires.Format(time.RFC3339))
				}
				var date []byte
				if date, err = file.Metadata.Expires.MarshalBinary(); err != nil {
					return err
				}
				tx.Bucket(s.eb).Put(file.Metadata.ID, date)
			}
			return nil
		})
	return file.Metadata.ID, err
}

// Get implements the Get method of the Filstorer interface
//
// It returns a store.File and a nil error if the file and the metadata is is found.
// If the filname for the given UUID is not found, a nil store.File and a CorruptedDataError is returned.
func (s *Bolt) Get(id uuid.UUID) (f *store.File, err error) {
	f = &store.File{Metadata: store.MetaData{}}
	err = s.db.View(
		func(tx *bolt.Tx) error {

			b := tx.Bucket(s.fb).Get(id)

			if len(b) < 1 {
				return store.FileNotFoundError("No file with ID " + id.String())
			}

			if name := tx.Bucket(s.nb).Get(id); name != nil {
				f.Metadata.Name = string(name)
			} else {
				return store.CorruptedDataError("No filename for ID " + id.String())
			}

			if date := tx.Bucket(s.eb).Get(id); date != nil {
				if err = f.Metadata.Expires.UnmarshalBinary(date); err != nil {
					return err
				}
			}

			if public := tx.Bucket(s.sb).Get(id); public != nil {
				f.Metadata.Public = true
			}

			// Copy and allocate as late as possible
			bu := make([]byte, len(b), cap(b))
			copy(bu, b)
			f.Content = bytes.NewBuffer(bu)

			return nil
		})

	if err != nil {
		return nil, err
	}

	return
}

// Remove implements the Remove method of the Filstorer interface.
func (s *Bolt) Remove(id uuid.UUID) (err error) {
	return s.db.Update(
		func(tx *bolt.Tx) error {

			for _, b := range [][]byte{s.eb, s.fb, s.nb, s.sb} {
				if err = tx.Bucket(b).Delete(id); err != nil {
					return err
				}
			}
			return nil
		})
}

func (s *Bolt) IsExpired(id uuid.UUID) (expired bool) {

	s.db.View(
		func(tx *bolt.Tx) error {
			expired = tx.Bucket(s.eb).Get(id) != nil
			return nil
		})
	return
}

func (s *Bolt) IsPublic(id uuid.UUID) (public bool) {
	s.db.View(
		func(tx *bolt.Tx) error {
			public = tx.Bucket(s.eb).Get(id) != nil
			return nil
		})
	return
}

func (s *Bolt) List() (list []store.MetaData) {
	s.db.View(
		func(tx *bolt.Tx) error {
			return tx.Bucket(s.nb).ForEach(
				func(k, v []byte) error {
					list = append(list, store.MetaData{ID: k, Name: string(v)})
					return nil
				})
		})
	return
}
