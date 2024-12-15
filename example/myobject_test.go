package example

import (
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/rogozhka/go-generate-mockgen/example/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMyObject_CreateItem(t *testing.T) {
	is := assert.New(t)

	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	idGenerator := mocks.NewMockUniqueIdentificatorGenerator(ctrlr)
	o := New(idGenerator)

	// randomID preparation
	testID, err := uuid.NewV7()
	is.Nil(err)

	idGenerator.EXPECT().GenerateID().Return(testID.String(), nil)

	res, err := o.CreateItem()
	is.Nil(err)

	// expect method returned the result of the GenerateID()
	is.Equal(testID.String(), res)
}

func TestMyObject_CreateItem_TableStyle(t *testing.T) {
	is := assert.New(t)

	ctrlr := gomock.NewController(t)
	defer ctrlr.Finish()

	type testCase struct {
		name       string
		inst       *myObject
		exp        string
		wantErr    bool
		expErrText string
	}

	cases := []testCase{
		func() testCase {
			const name = "simple positive case"

			idGenerator := mocks.NewMockUniqueIdentificatorGenerator(ctrlr)
			o := New(idGenerator)

			// randomID preparation
			testID, err := uuid.NewV7()
			is.Nil(err)

			idGenerator.EXPECT().GenerateID().Return(testID.String(), nil)

			return testCase{
				name:    name,
				inst:    o,
				exp:     testID.String(),
				wantErr: false,
			}
		}(),
		func() testCase {
			const name = "failed to generate id"

			idGenerator := mocks.NewMockUniqueIdentificatorGenerator(ctrlr)
			o := New(idGenerator)

			idGenerator.EXPECT().GenerateID().Return("", errors.New("no new numbers available"))

			return testCase{
				name:       name,
				inst:       o,
				exp:        "",
				wantErr:    true,
				expErrText: "generate ID: no new numbers available",
			}
		}(),
		func() testCase {
			const name = "nil instance"
			// just test there would be no panic on nil receiver

			return testCase{
				name:       name,
				inst:       nil,
				exp:        "",
				wantErr:    true,
				expErrText: "receiver is nil",
			}
		}(),
	}

	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.inst.CreateItem()
			if !tc.wantErr {
				is.Nil(err)
			} else {
				is.NotNil(err)
				is.Equal(tc.expErrText, err.Error())
			}
			is.Equal(tc.exp, res)
		})
	}
}
