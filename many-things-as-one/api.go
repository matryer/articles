package api

type Saver interface {
	Save() error
}

type Record struct {
	Bucket string
	Data   map[string]interface{}
}

func (r *Record) Save() error {
	// TODO: save one record
	return nil
}

type Savers []Saver

// Save saves all Saver objects.
func (r Savers) Save() error {
	for _, rec := range r {
		if err := rec.Save(); err != nil {
			return err
		}
	}
	return nil
}

type SaverFunc func() error

func (fn SaverFunc) Save() error {
	return fn()
}

func x() {
	records := Savers{&Record{}, &Record{}, &Record{}}
	records.Save()
}
