package envriron

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		parser := func(s string) (int, error) {
			return 1, nil
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.Nil(t, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})

	t.Run("Test Case 2", func(t *testing.T) {
		parser := func(s string) int {
			return 1
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.Nil(t, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})
}

func TestNilParser(t *testing.T) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
			assert.Equal(t, NilParser, err)

			wg.Done()
		}()

		UseParser(nil)
	}()

	wg.Wait()
}

func TestInvalidParserArgument(t *testing.T) {
	parser := 3

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			assert.NotNil(t, err)

			wg.Done()
		}()

		UseParser(parser)
	}()

	wg.Wait()
}

func TestInvalidParserFunctionArgument(t *testing.T) {
	parser := func(i int) (int, error) {
		return 0, nil
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
			assert.Equal(t, InvalidArgument, err)

			wg.Done()
		}()

		UseParser(parser)
	}()

	wg.Wait()
}

func TestInvalidParserFunctionReturnType(t *testing.T) {
	t.Run("Test Case 1", func(t *testing.T) {
		parser := func(s string) (int, int, int) {
			return 0, 0, 0
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				assert.Equal(t, InvalidReturnType, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})

	t.Run("Test Case 2", func(t *testing.T) {
		parser := func(s string) (int, int, int) {
			return 0, 0, 0
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				assert.Equal(t, InvalidReturnType, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})

	t.Run("Test Case 3", func(t *testing.T) {
		parser := func(s string) error {
			return nil
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				assert.Equal(t, InvalidReturnType, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})

	t.Run("Test Case 4", func(t *testing.T) {
		parser := func(s string) error {
			return nil
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				assert.Equal(t, InvalidReturnType, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})

	t.Run("Test Case 5", func(t *testing.T) {
		parser := func(s string) (int, int) {
			return 1, 1
		}

		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				assert.Equal(t, InvalidReturnType, err)

				wg.Done()
			}()

			UseParser(parser)
		}()

		wg.Wait()
	})
}
