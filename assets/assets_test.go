//  This file is part of the eliona project.
//  Copyright © 2022 LEICOM iTEC AG. All Rights Reserved.
//  ______ _ _
// |  ____| (_)
// | |__  | |_  ___  _ __   __ _
// |  __| | | |/ _ \| '_ \ / _` |
// | |____| | | (_) | | | | (_| |
// |______|_|_|\___/|_| |_|\__,_|
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED  TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
//  DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package assets

import (
	"database/sql"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpsertAssetType(t *testing.T) {
	mock := connectionMock()
	mock.ExpectExec("insert into public.asset_type").
		WithArgs("new_type", true, sql.NullString{String: "ITEC AG", Valid: true}, pgxmock.AnyArg(), sql.NullString{}, sql.NullString{}).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	err := UpsertAssetType(mock, AssetType{Id: "new_type", Custom: true, Vendor: "ITEC AG", Translation: Translation{German: "Neuer Typ", English: "New type"}})
	assert.Nil(t, err)
}

func TestUpsertAssetTypeAttribute(t *testing.T) {
	mock := connectionMock()
	mock.ExpectExec("insert into public.attribute_schema").
		WithArgs("new_type", "weather", "temperature", StatusSubtype, true, Translation{German: "Temperatur", English: "Temperature"}, sql.NullString{}).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	err := UpsertAssetTypeAttribute(mock, AssetTypeAttribute{AssetTypeId: "new_type", AttributeType: "weather", Id: "temperature", Subtype: StatusSubtype, Enable: true, Translation: Translation{German: "Temperatur", English: "Temperature"}})
	assert.Nil(t, err)
}

func TestUpsertAsset(t *testing.T) {
	mock := connectionMock()
	rows := mock.NewRows([]string{"asset_id"}).AddRow(4711)
	mock.ExpectQuery("insert into public.asset").
		WillReturnRows(rows)
	id := UpsertAsset(mock, Asset{ProjectId: "99", AssetTypeId: "new_type", GlobalAssetIdentifier: "dddddd", Description: "erd", Tags: []string{"a", "b"}})
	assert.Equal(t, 4711, id)
}

func TestGetAssetType(t *testing.T) {
	mock := connectionMock()

	rows := mock.NewRows([]string{"asset_type", "custom", "vendor", "translation", "urldoc", "icon"}).
		AddRow("new_type", true, "ITEC AG", Translation{German: "Neuer Typ", English: "New type"}, "", "")
	mock.ExpectQuery("select.+from public.asset_type").WillReturnRows(rows)
	assetType, err := GetAssetType(mock, "new_type")
	assert.Nil(t, err)
	assert.NotNil(t, assetType)

	empty := mock.NewRows([]string{"asset_type", "custom", "vendor", "translation", "urldoc", "icon"})
	mock.ExpectQuery("select.+from public.asset_type").WillReturnRows(empty)
	assetType, err = GetAssetType(mock, "foo bar")
	assert.Equal(t, err.Error(), "asset type foo bar not found")
}
