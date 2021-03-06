package mykad

// Copy pasta'd from https://en.wikipedia.org/wiki/Malaysian_identity_card#MyID. Delicious!!
var regions = map[placeOfBirthCode]string{
	1:  "Johor",
	2:  "Kedah",
	3:  "Kelantan",
	4:  "Malacca",
	5:  "Negeri Sembilan",
	6:  "Pahang",
	7:  "Penang",
	8:  "Perak",
	9:  "Perlis",
	10: "Selangor",
	11: "Terengganu",
	12: "Sabah",
	13: "Sarawak",
	14: "Federal Territory of Kuala Lumpur",
	15: "Federal Territory of Labuan",
	16: "Federal Territory of Putrajaya",
	21: "Johor",
	22: "Johor",
	23: "Johor",
	24: "Johor",
	25: "Kedah",
	26: "Kedah",
	27: "Kedah",
	28: "Kelantan",
	29: "Kelantan",
	30: "Malacca",
	31: "Negeri Sembilan",
	32: "Pahang",
	33: "Pahang",
	34: "Penang",
	35: "Penang",
	36: "Perak",
	37: "Perak",
	38: "Perak",
	39: "Perak",
	40: "Perlis",
	41: "Selangor",
	42: "Selangor",
	43: "Selangor",
	44: "Selangor",
	45: "Terengganu",
	46: "Terengganu",
	47: "Sabah",
	48: "Sabah",
	49: "Sabah",
	50: "Sarawak",
	51: "Sarawak",
	52: "Sarawak",
	53: "Sarawak",
	54: "Federal Territory of Kuala Lumpur",
	55: "Federal Territory of Kuala Lumpur",
	56: "Federal Territory of Kuala Lumpur",
	57: "Federal Territory of Kuala Lumpur",
	58: "Federal Territory of Labuan",
	59: "Negeri Sembilan",
	60: "Brunei",
	61: "Indonesia",
	62: "Cambodia / Democratic Kampuchea / Kampuchea",
	63: "Laos",
	64: "Myanmar",
	65: "Philippines",
	66: "Singapore",
	67: "Thailand",
	68: "Vietnam",
	71: `A person born outside Malaysia prior to 2001
Excluding those born abroad without holding High Quality Identity Card`,
	72: `A person born outside Malaysia prior to 2001
Excluding those born abroad without holding High Quality Identity Card`,
	74: "China",
	75: "India",
	76: "Pakistan",
	77: "Saudi Arabia",
	78: "Sri Lanka",
	79: "Bangladesh",
	82: "Unknown state",
	83: "American Samoa / Asia-Pacific / Australia / Christmas Island / Cocos (Keeling) Islands / Cook Islands / Fiji / French Polynesia / Guam / Heard Island and McDonald Islands / Marshall Islands / Micronesia / New Caledonia / New Zealand / Niue / Norfolk Island / Papua New Guinea / Timor Leste / Tokelau / United States Minor Outlying Islands / Wallis and Futuna Islands",
	84: "Anguilla / Argentina / Aruba / Bolivia / Brazil / Chile / Colombia / Ecuador / French Guinea / Guadeloupe / Guyana / Paraguay / Peru / South America / South Georgia and the South Sandwich Islands / Suriname / Uruguay / Venezuela",
	85: "Africa / Algeria / Angola / Botswana / Burundi / Cameroon / Central African Republic / Chad / Congo-Brazzaville / Congo-Kinshasa / Djibouti / Egypt / Eritrea / Ethiopia / Gabon / Gambia / Ghana / Guinea / Kenya / Liberia / Malawi / Mali / Mauritania / Mayotte / Morocco / Mozambique / Namibia / Niger / Nigeria / Rwanda / Réunion / Senegal / Sierra Leone / Somalia / South Africa / Sudan / Swaziland / Tanzania / Togo / Tonga / Tunisia / Uganda / Western Sahara / Zaire / Zambia / Zimbabwe",
	86: "Armenia / Austria / Belgium / Cyprus / Denmark / Europe / Faroe Islands / France / Finland / Finland, Metropolitan / Germany / Germany, Democratic Republic / Germany, Federal Republic / Greece / Holy See (Vatican City) / Italy / Luxembourg / Macedonia / Malta / Mediterranean / Monaco / Netherlands / Norway / Portugal / Republic of Moldova / Slovakia / Slovenia / Spain / Sweden / Switzerland / United Kingdom-Dependent Territories / United Kingdom-National Overseas / United Kingdom-Overseas Citizen / United Kingdom-Protected Person / United Kingdom-Subject",
	87: "Britain / Great Britain / Ireland",
	88: "Bahrain / Iran / Iraq / Palestine / Jordan / Kuwait / Lebanon / Middle East / Oman / Qatar / Republic of Yemen / Syria / Turkey / United Arab Emirates / Yemen Arab Republic / Yemen People's Democratic Republic / Israel",
	89: "Far East / Japan / North Korea / South Korea / Taiwan",
	90: "Bahamas / Barbados / Belize / Caribbean / Costa Rica / Cuba / Dominica / Dominican Republic / El Salvador / Grenada / Guatemala / Haiti / Honduras / Jamaica / Martinique / Mexico / Nicaragua / Panama / Puerto Rico / Saint Kitts and Nevis / Saint Lucia / Saint Vincent and the Grenadines / Trinidad and Tobago / Turks and Caicos Islands / Virgin Islands (USA)",
	91: "Canada / Greenland / Netherlands Antilles / North America / Saint Pierre and Miquelon / United States of America",
	92: "Albania / Belarus / Bosnia and Herzegovina / Bulgaria / Byelorussia / Croatia / Czech Republic / Czechoslovakia / Estonia / Georgia / Hungary / Latvia / Lithuania / Montenegro / Poland / Republic of Kosovo / Romania / Russian Federation / Serbia / Soviet Union / U.S.S.R. / Ukraine",
	93: "Afghanistan / Andorra / Antarctica / Antigua and Barbuda / Azerbaijan / Benin / Bermuda / Bhutan / Bora Bora / Bouvet Island / British Indian Ocean Territory / Burkina Faso / Cape Verde / Cayman Islands / Comoros / Dahomey / Equatorial Guinea / Falkland Islands / French Southern Territories / Gibraltar / Guinea-Bissau / Hong Kong / Iceland / Ivory Coast / Kazakhstan / Kiribati / Kyrgyzstan / Lesotho / Libya / Liechtenstein / Macau / Madagascar / Maghribi / Malagasy / Maldives / Mauritius / Mongolia / Montserrat / Nauru / Nepal / Northern Marianas Islands / Outer Mongolia / Palau / Palestine / Pitcairn Islands / Saint Helena / Saint Lucia / Saint Vincent and the Grenadines / Samoa / San Marino / São Tomé and Príncipe / Seychelles / Solomon Islands / Svalbard and Jan Mayen Islands / Tajikistan / Turkmenistan / Tuvalu / Upper Volta / Uzbekistan / Vanuatu / Vatican City / Virgin Islands (British) / Western Samoa / Yugoslavia",
	98: "Stateless / Stateless Person Article 1/1954",
	99: "Mecca / Neutral Zone / No Information / Refugee / Refugee Article 1/1951 / United Nations Specialized Agency / United Nations Organization / Unspecified Nationality",
}
