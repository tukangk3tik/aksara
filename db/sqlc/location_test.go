package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const provinceName = "Nusa"
const provinceID = 53
const regencyName = "Kota Kupang"
const regencyID = 5371

func TestLocationProvince(t *testing.T) {
	provinces, err := testQueries.LocationProvince(context.Background(), &LocationProvinceParams{
		Name:   fmt.Sprintf("%%%s%%", provinceName),
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(provinces), 2)
	require.NotEmpty(t, provinces)

	for _, province := range provinces {
		require.NotEmpty(t, province)
	}
}

func TestLocationRegencyByProvince(t *testing.T) {
	regencies, err := testQueries.LocationRegencyByProvince(context.Background(), &LocationRegencyByProvinceParams{
		ProvinceID: provinceID,
		Name:       fmt.Sprintf("%%%s%%", regencyName),
		Limit:      10,
		Offset:     0,
	})
	require.NoError(t, err)
	require.Equal(t, len(regencies), 1)
	require.NotEmpty(t, regencies)

	for _, regency := range regencies {
		require.NotEmpty(t, regency)
	}
}

func TestLocationDistrictByRegency(t *testing.T) {
	districts, err := testQueries.LocationDistrictByRegency(context.Background(), &LocationDistrictByRegencyParams{
		RegencyID: regencyID,
		Name:      fmt.Sprintf("%%%s%%", "Oebobo"),
		Limit:     10,
		Offset:    0,
	})
	require.NoError(t, err)
	require.Equal(t, len(districts), 1)
	require.NotEmpty(t, districts)

	for _, district := range districts {
		require.NotEmpty(t, district)
	}
}

