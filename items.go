package postman

import "errors"

// Items are the basic unit for a Postman collection.
// It can either be a request (Item) or a folder (ItemGroup).
type Items interface {
	IsGroup() bool
}

func createItemCollection(items []interface{}) (itemCollection []Items, err error) {
	for _, i := range items {
		item, err := createItemFromInterface(i)

		if err != nil {
			return nil, err
		}

		itemCollection = append(itemCollection, item)
	}

	return itemCollection, nil
}

func createItemFromInterface(i interface{}) (item Items, err error) {
	dict, ok := i.(map[string]interface{})

	if !ok {
		return nil, errors.New("Unsupported interface")
	}

	if _, found := dict["item"]; found {
		item, err = decodeItemGroup(dict)
	} else {
		item, err = decodeItem(dict)
	}

	return
}
