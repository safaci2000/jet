package postgres

const testMultipleNestedAliasedSlices = `
[
	{
		"EmployeeId": 1,
		"LastName": "Adams",
		"FirstName": "Andrew",
		"Title": "General Manager",
		"ReportsTo": null,
		"BirthDate": "1962-02-18T00:00:00Z",
		"HireDate": "2002-08-14T00:00:00Z",
		"Address": "11120 Jasper Ave NW",
		"City": "Edmonton",
		"State": "AB",
		"Country": "Canada",
		"PostalCode": "T5K 2N1",
		"Phone": "+1 (780) 428-9482",
		"Fax": "+1 (780) 428-3457",
		"Email": "andrew@chinookcorp.com",
		"Employees": [
			{
				"EmployeeId": 2,
				"LastName": "Edwards",
				"FirstName": "Nancy",
				"Title": "Sales Manager",
				"ReportsTo": 1,
				"BirthDate": "1958-12-08T00:00:00Z",
				"HireDate": "2002-05-01T00:00:00Z",
				"Address": "825 8 Ave SW",
				"City": "Calgary",
				"State": "AB",
				"Country": "Canada",
				"PostalCode": "T2P 2T3",
				"Phone": "+1 (403) 262-3443",
				"Fax": "+1 (403) 262-3322",
				"Email": "nancy@chinookcorp.com"
			},
			{
				"EmployeeId": 6,
				"LastName": "Mitchell",
				"FirstName": "Michael",
				"Title": "IT Manager",
				"ReportsTo": 1,
				"BirthDate": "1973-07-01T00:00:00Z",
				"HireDate": "2003-10-17T00:00:00Z",
				"Address": "5827 Bowness Road NW",
				"City": "Calgary",
				"State": "AB",
				"Country": "Canada",
				"PostalCode": "T3B 0C5",
				"Phone": "+1 (403) 246-9887",
				"Fax": "+1 (403) 246-9899",
				"Email": "michael@chinookcorp.com"
			}
		],
		"EmployeesCustom1": [
			{
				"EmployeeId": 2,
				"LastName": "Edwards",
				"FirstName": "Nancy",
				"Title": "Sales Manager",
				"ReportsTo": 1,
				"BirthDate": "1958-12-08T00:00:00Z",
				"HireDate": "2002-05-01T00:00:00Z",
				"Address": "825 8 Ave SW",
				"City": "Calgary",
				"State": "AB",
				"Country": "Canada",
				"PostalCode": "T2P 2T3",
				"Phone": "+1 (403) 262-3443",
				"Fax": "+1 (403) 262-3322",
				"Email": "nancy@chinookcorp.com",
				"Employees1": [
					{
						"EmployeeId": 3,
						"LastName": "Peacock",
						"FirstName": "Jane",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1973-08-29T00:00:00Z",
						"HireDate": "2002-04-01T00:00:00Z",
						"Address": "1111 6 Ave SW",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T2P 5M5",
						"Phone": "+1 (403) 262-3443",
						"Fax": "+1 (403) 262-6712",
						"Email": "jane@chinookcorp.com"
					},
					{
						"EmployeeId": 4,
						"LastName": "Park",
						"FirstName": "Margaret",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1947-09-19T00:00:00Z",
						"HireDate": "2003-05-03T00:00:00Z",
						"Address": "683 10 Street SW",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T2P 5G3",
						"Phone": "+1 (403) 263-4423",
						"Fax": "+1 (403) 263-4289",
						"Email": "margaret@chinookcorp.com"
					},
					{
						"EmployeeId": 5,
						"LastName": "Johnson",
						"FirstName": "Steve",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1965-03-03T00:00:00Z",
						"HireDate": "2003-10-17T00:00:00Z",
						"Address": "7727B 41 Ave",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T3B 1Y7",
						"Phone": "1 (780) 836-9987",
						"Fax": "1 (780) 836-9543",
						"Email": "steve@chinookcorp.com"
					}
				],
				"EmployeesCustom2": [
					{
						"EmployeeId": 3,
						"LastName": "Peacock",
						"FirstName": "Jane",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1973-08-29T00:00:00Z",
						"HireDate": "2002-04-01T00:00:00Z",
						"Address": "1111 6 Ave SW",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T2P 5M5",
						"Phone": "+1 (403) 262-3443",
						"Fax": "+1 (403) 262-6712",
						"Email": "jane@chinookcorp.com",
						"Customers2": [
							{
								"CustomerId": 37,
								"FirstName": "Fynn",
								"LastName": "Zimmermann",
								"Company": null,
								"Address": "Berger Stra�e 10",
								"City": "Frankfurt",
								"State": null,
								"Country": "Germany",
								"PostalCode": "60316",
								"Phone": "+49 069 40598889",
								"Fax": null,
								"Email": "fzimmermann@yahoo.de",
								"SupportRepId": 3
							},
							{
								"CustomerId": 38,
								"FirstName": "Niklas",
								"LastName": "Schr�der",
								"Company": null,
								"Address": "Barbarossastra�e 19",
								"City": "Berlin",
								"State": null,
								"Country": "Germany",
								"PostalCode": "10779",
								"Phone": "+49 030 2141444",
								"Fax": null,
								"Email": "nschroder@surfeu.de",
								"SupportRepId": 3
							},
							{
								"CustomerId": 42,
								"FirstName": "Wyatt",
								"LastName": "Girard",
								"Company": null,
								"Address": "9, Place Louis Barthou",
								"City": "Bordeaux",
								"State": null,
								"Country": "France",
								"PostalCode": "33000",
								"Phone": "+33 05 56 96 96 96",
								"Fax": null,
								"Email": "wyatt.girard@yahoo.fr",
								"SupportRepId": 3
							}
						]
					},
					{
						"EmployeeId": 4,
						"LastName": "Park",
						"FirstName": "Margaret",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1947-09-19T00:00:00Z",
						"HireDate": "2003-05-03T00:00:00Z",
						"Address": "683 10 Street SW",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T2P 5G3",
						"Phone": "+1 (403) 263-4423",
						"Fax": "+1 (403) 263-4289",
						"Email": "margaret@chinookcorp.com",
						"Customers2": [
							{
								"CustomerId": 39,
								"FirstName": "Camille",
								"LastName": "Bernard",
								"Company": null,
								"Address": "4, Rue Milton",
								"City": "Paris",
								"State": null,
								"Country": "France",
								"PostalCode": "75009",
								"Phone": "+33 01 49 70 65 65",
								"Fax": null,
								"Email": "camille.bernard@yahoo.fr",
								"SupportRepId": 4
							},
							{
								"CustomerId": 40,
								"FirstName": "Dominique",
								"LastName": "Lefebvre",
								"Company": null,
								"Address": "8, Rue Hanovre",
								"City": "Paris",
								"State": null,
								"Country": "France",
								"PostalCode": "75002",
								"Phone": "+33 01 47 42 71 71",
								"Fax": null,
								"Email": "dominiquelefebvre@gmail.com",
								"SupportRepId": 4
							}
						]
					},
					{
						"EmployeeId": 5,
						"LastName": "Johnson",
						"FirstName": "Steve",
						"Title": "Sales Support Agent",
						"ReportsTo": 2,
						"BirthDate": "1965-03-03T00:00:00Z",
						"HireDate": "2003-10-17T00:00:00Z",
						"Address": "7727B 41 Ave",
						"City": "Calgary",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T3B 1Y7",
						"Phone": "1 (780) 836-9987",
						"Fax": "1 (780) 836-9543",
						"Email": "steve@chinookcorp.com",
						"Customers2": [
							{
								"CustomerId": 41,
								"FirstName": "Marc",
								"LastName": "Dubois",
								"Company": null,
								"Address": "11, Place Bellecour",
								"City": "Lyon",
								"State": null,
								"Country": "France",
								"PostalCode": "69002",
								"Phone": "+33 04 78 30 30 30",
								"Fax": null,
								"Email": "marc.dubois@hotmail.com",
								"SupportRepId": 5
							}
						]
					}
				]
			},
			{
				"EmployeeId": 6,
				"LastName": "Mitchell",
				"FirstName": "Michael",
				"Title": "IT Manager",
				"ReportsTo": 1,
				"BirthDate": "1973-07-01T00:00:00Z",
				"HireDate": "2003-10-17T00:00:00Z",
				"Address": "5827 Bowness Road NW",
				"City": "Calgary",
				"State": "AB",
				"Country": "Canada",
				"PostalCode": "T3B 0C5",
				"Phone": "+1 (403) 246-9887",
				"Fax": "+1 (403) 246-9899",
				"Email": "michael@chinookcorp.com",
				"Employees1": [
					{
						"EmployeeId": 7,
						"LastName": "King",
						"FirstName": "Robert",
						"Title": "IT Staff",
						"ReportsTo": 6,
						"BirthDate": "1970-05-29T00:00:00Z",
						"HireDate": "2004-01-02T00:00:00Z",
						"Address": "590 Columbia Boulevard West",
						"City": "Lethbridge",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T1K 5N8",
						"Phone": "+1 (403) 456-9986",
						"Fax": "+1 (403) 456-8485",
						"Email": "robert@chinookcorp.com"
					},
					{
						"EmployeeId": 8,
						"LastName": "Callahan",
						"FirstName": "Laura",
						"Title": "IT Staff",
						"ReportsTo": 6,
						"BirthDate": "1968-01-09T00:00:00Z",
						"HireDate": "2004-03-04T00:00:00Z",
						"Address": "923 7 ST NW",
						"City": "Lethbridge",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T1H 1Y8",
						"Phone": "+1 (403) 467-3351",
						"Fax": "+1 (403) 467-8772",
						"Email": "laura@chinookcorp.com"
					}
				],
				"EmployeesCustom2": [
					{
						"EmployeeId": 7,
						"LastName": "King",
						"FirstName": "Robert",
						"Title": "IT Staff",
						"ReportsTo": 6,
						"BirthDate": "1970-05-29T00:00:00Z",
						"HireDate": "2004-01-02T00:00:00Z",
						"Address": "590 Columbia Boulevard West",
						"City": "Lethbridge",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T1K 5N8",
						"Phone": "+1 (403) 456-9986",
						"Fax": "+1 (403) 456-8485",
						"Email": "robert@chinookcorp.com",
						"Customers2": null
					},
					{
						"EmployeeId": 8,
						"LastName": "Callahan",
						"FirstName": "Laura",
						"Title": "IT Staff",
						"ReportsTo": 6,
						"BirthDate": "1968-01-09T00:00:00Z",
						"HireDate": "2004-03-04T00:00:00Z",
						"Address": "923 7 ST NW",
						"City": "Lethbridge",
						"State": "AB",
						"Country": "Canada",
						"PostalCode": "T1H 1Y8",
						"Phone": "+1 (403) 467-3351",
						"Fax": "+1 (403) 467-8772",
						"Email": "laura@chinookcorp.com",
						"Customers2": null
					}
				]
			}
		]
	}
]
`