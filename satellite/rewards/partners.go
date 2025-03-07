// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package rewards

// DefaultPartnersDB is current default settings.
var DefaultPartnersDB = func() PartnersDB {
	list := DefaultPartners()
	db, err := NewPartnersStaticDB(&list)
	if err != nil {
		panic(err)
	}
	return db
}()

// DefaultPartners lists Storj default open-source partners.
func DefaultPartners() PartnerList {
	return PartnerList{
		Partners: []PartnerInfo{{
			Name: "Blocknify",
			ID:   "120bf202-8252-437e-ac12-0e364bee852e",
		}, {
			Name: "Breaker",
			ID:   "53688ea5-8695-4060-a2c6-b56969217909",
		}, {
			Name: "Confluent",
			ID:   "2fb801c6-a6d7-4d82-a838-32fef98cc398",
		}, {
			Name: "Consensys",
			ID:   "e28c8847-b323-4a7d-8111-25a0578a58bb",
		}, {
			Name: "Couchbase",
			ID:   "0af89ac1-0189-42c6-a47c-e169780b3818",
		}, {
			Name: "Digital Ocean",
			ID:   "881b92f6-77aa-42ee-961a-b80009d45dd8",
		}, {
			Name: "Deloitte",
			ID:   "cadac3fb-6a3f-4d17-9748-cc66d0617d55",
		}, {
			Name: "DVLabs",
			ID:   "53fb82d7-73ff-4a1a-ab0c-6968cffc850e",
		}, {
			Name: "Fluree",
			ID:   "86c33256-cded-434c-aaac-405343974394",
		}, {
			Name: "Flexential",
			ID:   "3e1b911a-c778-47ea-878c-9f3f264f8bc1",
		}, {
			Name: "Heroku",
			ID:   "706011f3-400e-45eb-a796-90cce2a7d67e",
		}, {
			Name: "Infura",
			ID:   "1519bdee-ed18-45fe-86c6-4c7fa9668a14",
		}, {
			Name: "GroundX",
			ID:   "e56c6a65-d5bf-457a-a414-e55c36624f73",
		}, {
			Name: "MariaDB",
			ID:   "8ee019ef-2aae-4867-9c18-41c65ea318c4",
		}, {
			Name: "MongoDB",
			ID:   "bbd340b2-0ae4-4254-af90-eaba6c273abb",
		}, {
			Name: "Netki",
			ID:   "3405a882-0cb2-4f91-a6e0-21be193b80e5",
		}, {
			Name: "FileZilla",
			ID:   "a1ba07a4-e095-4a43-914c-1d56c9ff5afd",
		}, {
			Name: "InfluxDB",
			ID:   "e50a17b3-4d82-4da7-8719-09312a83685d",
		}, {
			Name: "Mysterium Network",
			ID:   "c10228c2-af70-4e4d-be49-e8bfbe9ca8ef",
		}, {
			Name: "Kafka",
			ID:   "OSPP005",
		}, {
			Name: "Minio",
			ID:   "5bffe844-5da7-4aa9-bf37-7d695cf819f2",
		}, {
			Name: "Nextcloud",
			ID:   "42f588fb-f39d-4886-81af-b614ca16ce37",
		}, {
			Name: "Node Haven",
			ID:   "3b53a9b3-2005-476c-9ffd-894ed832abe4",
		}, {
			Name: "Plesk",
			ID:   "dc01ed96-2990-4819-9cb3-45d4846b9ad1",
		}, {
			Name: "Pydio",
			ID:   "b02b9f0d-fac7-439c-8ba2-0c4634d5826f",
		}, {
			Name: "Raiden Network",
			ID:   "57855387-5a58-4a2b-97d2-15b1d76eea3c",
		}, {
			Name: "Satoshi Soup",
			ID:   "4400d796-3777-4964-8536-22a4ae439ed3",
		}, {
			Name: "Sirin Labs",
			ID:   "6e40f882-ef77-4a5d-b5ad-18525d3df023",
		}, {
			Name: "Status Messenger",
			ID:   "b6114126-c06d-49f9-8d23-3e0dd2e350ab",
		}, {
			Name: "Temporal",
			ID:   "aeedbe32-1519-4320-b2f4-33725c65af54",
		}, {
			Name: "Terminal.co",
			ID:   "7bf23e53-6393-4bd0-8bf9-53ecf0de742f",
		}, {
			Name: "Zenko",
			ID:   "8cd605fa-ad00-45b6-823e-550eddc611d6",
		}},
	}
}
