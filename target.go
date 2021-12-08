package fz

// must provide these functions to control target

type Start func() error

type Stop func() error

type Do func(action Action) error
