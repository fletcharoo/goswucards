package goswucards

import (
	"embed"
)

//go:embed images/*
var ImagesFS embed.FS

type Card struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

var Cards = map[string]Card{
	"SOR001": {
		ID: "SOR001",
		Name: "Director Krennic",
	},
	"SOR002": {
		ID: "SOR002",
		Name: "Iden Versio",
	},
	"SOR003": {
		ID: "SOR003",
		Name: "Chewbacca",
	},
	"SOR004": {
		ID: "SOR004",
		Name: "Chirrut Îmwe",
	},
	"SOR005": {
		ID: "SOR005",
		Name: "Luke Skywalker",
	},
	"SOR006": {
		ID: "SOR006",
		Name: "Emperor Palpatine",
	},
	"SOR007": {
		ID: "SOR007",
		Name: "Grand Moff Tarkin",
	},
	"SOR008": {
		ID: "SOR008",
		Name: "Hera Syndulla",
	},
	"SOR009": {
		ID: "SOR009",
		Name: "Leia Organa",
	},
	"SOR010": {
		ID: "SOR010",
		Name: "Darth Vader",
	},
	"SOR011": {
		ID: "SOR011",
		Name: "Grand Inquisitor",
	},
	"SOR012": {
		ID: "SOR012",
		Name: "IG-88",
	},
	"SOR013": {
		ID: "SOR013",
		Name: "Cassian Andor",
	},
	"SOR014": {
		ID: "SOR014",
		Name: "Sabine Wren",
	},
	"SOR015": {
		ID: "SOR015",
		Name: "Boba Fett",
	},
	"SOR016": {
		ID: "SOR016",
		Name: "Grand Admiral Thrawn",
	},
	"SOR017": {
		ID: "SOR017",
		Name: "Han Solo",
	},
	"SOR018": {
		ID: "SOR018",
		Name: "Jyn Erso",
	},
	"SOR019": {
		ID: "SOR019",
		Name: "Security Complex",
	},
	"SOR020": {
		ID: "SOR020",
		Name: "Capital City",
	},
	"SOR021": {
		ID: "SOR021",
		Name: "Dagobah Swamp",
	},
	"SOR022": {
		ID: "SOR022",
		Name: "Energy Conversion Lab",
	},
	"SOR023": {
		ID: "SOR023",
		Name: "Command Center",
	},
	"SOR024": {
		ID: "SOR024",
		Name: "Echo Base",
	},
	"SOR025": {
		ID: "SOR025",
		Name: "Tarkintown",
	},
	"SOR026": {
		ID: "SOR026",
		Name: "Catacombs of Cadera",
	},
	"SOR027": {
		ID: "SOR027",
		Name: "Kestro City",
	},
	"SOR028": {
		ID: "SOR028",
		Name: "Jedha City",
	},
	"SOR029": {
		ID: "SOR029",
		Name: "Administrator's Tower",
	},
	"SOR030": {
		ID: "SOR030",
		Name: "Chopper Base",
	},
	"SOR031": {
		ID: "SOR031",
		Name: "Inferno Four",
	},
	"SOR032": {
		ID: "SOR032",
		Name: "Scout Bike Pursuer",
	},
	"SOR033": {
		ID: "SOR033",
		Name: "Death Trooper",
	},
	"SOR034": {
		ID: "SOR034",
		Name: "Del Meeko",
	},
	"SOR035": {
		ID: "SOR035",
		Name: "Lieutenant Childsen",
	},
	"SOR036": {
		ID: "SOR036",
		Name: "Gideon Hask",
	},
	"SOR037": {
		ID: "SOR037",
		Name: "Academy Defense Walker",
	},
	"SOR038": {
		ID: "SOR038",
		Name: "Count Dooku",
	},
	"SOR039": {
		ID: "SOR039",
		Name: "AT-AT Suppressor",
	},
	"SOR040": {
		ID: "SOR040",
		Name: "Avenger",
	},
	"SOR041": {
		ID: "SOR041",
		Name: "Power of the Dark Side",
	},
	"SOR042": {
		ID: "SOR042",
		Name: "Search Your Feelings",
	},
	"SOR043": {
		ID: "SOR043",
		Name: "Superlaser Blast",
	},
	"SOR044": {
		ID: "SOR044",
		Name: "Restored ARC-170",
	},
	"SOR045": {
		ID: "SOR045",
		Name: "Yoda",
	},
	"SOR046": {
		ID: "SOR046",
		Name: "Consular Security Force",
	},
	"SOR047": {
		ID: "SOR047",
		Name: "Kanan Jarrus",
	},
	"SOR048": {
		ID: "SOR048",
		Name: "Vigilant Honor Guards",
	},
	"SOR049": {
		ID: "SOR049",
		Name: "Obi-Wan Kenobi",
	},
	"SOR050": {
		ID: "SOR050",
		Name: "The Ghost",
	},
	"SOR051": {
		ID: "SOR051",
		Name: "Luke Skywalker",
	},
	"SOR052": {
		ID: "SOR052",
		Name: "Redemption",
	},
	"SOR053": {
		ID: "SOR053",
		Name: "Luke's Lightsaber",
	},
	"SOR054": {
		ID: "SOR054",
		Name: "Jedi Lightsaber",
	},
	"SOR055": {
		ID: "SOR055",
		Name: "The Force Is With Me",
	},
	"SOR056": {
		ID: "SOR056",
		Name: "Bendu",
	},
	"SOR057": {
		ID: "SOR057",
		Name: "Protector",
	},
	"SOR058": {
		ID: "SOR058",
		Name: "Vigilance",
	},
	"SOR059": {
		ID: "SOR059",
		Name: "2-1B Surgical Droid",
	},
	"SOR060": {
		ID: "SOR060",
		Name: "Distant Patroller",
	},
	"SOR061": {
		ID: "SOR061",
		Name: "Guardian of the Whills",
	},
	"SOR062": {
		ID: "SOR062",
		Name: "Regional Governor",
	},
	"SOR063": {
		ID: "SOR063",
		Name: "Cloud City Wing Guard",
	},
	"SOR064": {
		ID: "SOR064",
		Name: "Wilderness Fighter",
	},
	"SOR065": {
		ID: "SOR065",
		Name: "Baze Malbus",
	},
	"SOR066": {
		ID: "SOR066",
		Name: "System Patrol Craft",
	},
	"SOR067": {
		ID: "SOR067",
		Name: "Rugged Survivors",
	},
	"SOR068": {
		ID: "SOR068",
		Name: "Cargo Juggernaut",
	},
	"SOR069": {
		ID: "SOR069",
		Name: "Resilient",
	},
	"SOR070": {
		ID: "SOR070",
		Name: "Devotion",
	},
	"SOR071": {
		ID: "SOR071",
		Name: "Electrostaff",
	},
	"SOR072": {
		ID: "SOR072",
		Name: "Entrenched",
	},
	"SOR073": {
		ID: "SOR073",
		Name: "Moment of Peace",
	},
	"SOR074": {
		ID: "SOR074",
		Name: "Repair",
	},
	"SOR075": {
		ID: "SOR075",
		Name: "It Binds All Things",
	},
	"SOR076": {
		ID: "SOR076",
		Name: "Make an Opening",
	},
	"SOR077": {
		ID: "SOR077",
		Name: "Takedown",
	},
	"SOR078": {
		ID: "SOR078",
		Name: "Vanquish",
	},
	"SOR079": {
		ID: "SOR079",
		Name: "Admiral Piett",
	},
	"SOR080": {
		ID: "SOR080",
		Name: "General Tagge",
	},
	"SOR081": {
		ID: "SOR081",
		Name: "Seasoned Shoretrooper",
	},
	"SOR082": {
		ID: "SOR082",
		Name: "Emperor's Royal Guard",
	},
	"SOR083": {
		ID: "SOR083",
		Name: "Superlaser Technician",
	},
	"SOR084": {
		ID: "SOR084",
		Name: "Grand Moff Tarkin",
	},
	"SOR085": {
		ID: "SOR085",
		Name: "Rukh",
	},
	"SOR086": {
		ID: "SOR086",
		Name: "Gladiator Star Destroyer",
	},
	"SOR087": {
		ID: "SOR087",
		Name: "Darth Vader",
	},
	"SOR088": {
		ID: "SOR088",
		Name: "Blizzard Assault AT-AT",
	},
	"SOR089": {
		ID: "SOR089",
		Name: "Relentless",
	},
	"SOR090": {
		ID: "SOR090",
		Name: "Devastator",
	},
	"SOR091": {
		ID: "SOR091",
		Name: "The Emperor's Legion",
	},
	"SOR092": {
		ID: "SOR092",
		Name: "Overwhelming Barrage",
	},
	"SOR093": {
		ID: "SOR093",
		Name: "Alliance Dispatcher",
	},
	"SOR094": {
		ID: "SOR094",
		Name: "Bail Organa",
	},
	"SOR095": {
		ID: "SOR095",
		Name: "Battlefield Marine",
	},
	"SOR096": {
		ID: "SOR096",
		Name: "Mon Mothma",
	},
	"SOR097": {
		ID: "SOR097",
		Name: "Admiral Ackbar",
	},
	"SOR098": {
		ID: "SOR098",
		Name: "Echo Base Defender",
	},
	"SOR099": {
		ID: "SOR099",
		Name: "Bright Hope",
	},
	"SOR100": {
		ID: "SOR100",
		Name: "Wedge Antilles",
	},
	"SOR101": {
		ID: "SOR101",
		Name: "Rogue Squadron Skirmisher",
	},
	"SOR102": {
		ID: "SOR102",
		Name: "Home One",
	},
	"SOR103": {
		ID: "SOR103",
		Name: "Rebel Assault",
	},
	"SOR104": {
		ID: "SOR104",
		Name: "U-Wing Reinforcement",
	},
	"SOR105": {
		ID: "SOR105",
		Name: "General Krell",
	},
	"SOR106": {
		ID: "SOR106",
		Name: "Attack Pattern Delta",
	},
	"SOR107": {
		ID: "SOR107",
		Name: "Command",
	},
	"SOR108": {
		ID: "SOR108",
		Name: "Vanguard Infantry",
	},
	"SOR109": {
		ID: "SOR109",
		Name: "Colonel Yularen",
	},
	"SOR110": {
		ID: "SOR110",
		Name: "Frontline Shuttle",
	},
	"SOR111": {
		ID: "SOR111",
		Name: "Patrolling V-Wing",
	},
	"SOR112": {
		ID: "SOR112",
		Name: "Consortium StarViper",
	},
	"SOR113": {
		ID: "SOR113",
		Name: "Homestead Militia",
	},
	"SOR114": {
		ID: "SOR114",
		Name: "Escort Skiff",
	},
	"SOR115": {
		ID: "SOR115",
		Name: "Agent Kallus",
	},
	"SOR116": {
		ID: "SOR116",
		Name: "Steadfast Battalion",
	},
	"SOR117": {
		ID: "SOR117",
		Name: "Mercenary Company",
	},
	"SOR118": {
		ID: "SOR118",
		Name: "97th Legion",
	},
	"SOR119": {
		ID: "SOR119",
		Name: "Reinforcement Walker",
	},
	"SOR120": {
		ID: "SOR120",
		Name: "Academy Training",
	},
	"SOR121": {
		ID: "SOR121",
		Name: "Hardpoint Heavy Blaster",
	},
	"SOR122": {
		ID: "SOR122",
		Name: "Traitorous",
	},
	"SOR123": {
		ID: "SOR123",
		Name: "Recruit",
	},
	"SOR124": {
		ID: "SOR124",
		Name: "Tactical Advantage",
	},
	"SOR125": {
		ID: "SOR125",
		Name: "Prepare for Takeoff",
	},
	"SOR126": {
		ID: "SOR126",
		Name: "Resupply",
	},
	"SOR127": {
		ID: "SOR127",
		Name: "Strike True",
	},
	"SOR128": {
		ID: "SOR128",
		Name: "Death Star Stormtrooper",
	},
	"SOR129": {
		ID: "SOR129",
		Name: "Admiral Ozzel",
	},
	"SOR130": {
		ID: "SOR130",
		Name: "First Legion Snowtrooper",
	},
	"SOR131": {
		ID: "SOR131",
		Name: "Fifth Brother",
	},
	"SOR132": {
		ID: "SOR132",
		Name: "Imperial Interceptor",
	},
	"SOR133": {
		ID: "SOR133",
		Name: "Seventh Sister",
	},
	"SOR134": {
		ID: "SOR134",
		Name: "Ruthless Raider",
	},
	"SOR135": {
		ID: "SOR135",
		Name: "Emperor Palpatine",
	},
	"SOR136": {
		ID: "SOR136",
		Name: "Vader's Lightsaber",
	},
	"SOR137": {
		ID: "SOR137",
		Name: "Fallen Lightsaber",
	},
	"SOR138": {
		ID: "SOR138",
		Name: "Force Lightning",
	},
	"SOR139": {
		ID: "SOR139",
		Name: "Force Choke",
	},
	"SOR140": {
		ID: "SOR140",
		Name: "SpecForce Soldier",
	},
	"SOR141": {
		ID: "SOR141",
		Name: "Green Squadron A-Wing",
	},
	"SOR142": {
		ID: "SOR142",
		Name: "Sabine Wren",
	},
	"SOR143": {
		ID: "SOR143",
		Name: "Fighters For Freedom",
	},
	"SOR144": {
		ID: "SOR144",
		Name: "Red Three",
	},
	"SOR145": {
		ID: "SOR145",
		Name: "K-2SO",
	},
	"SOR146": {
		ID: "SOR146",
		Name: "Zeb Orrelios",
	},
	"SOR147": {
		ID: "SOR147",
		Name: "Black One",
	},
	"SOR148": {
		ID: "SOR148",
		Name: "Guerilla Attack Pod",
	},
	"SOR149": {
		ID: "SOR149",
		Name: "Mace Windu",
	},
	"SOR150": {
		ID: "SOR150",
		Name: "Heroic Sacrifice",
	},
	"SOR151": {
		ID: "SOR151",
		Name: "Karabast",
	},
	"SOR152": {
		ID: "SOR152",
		Name: "For A Cause I Believe In",
	},
	"SOR153": {
		ID: "SOR153",
		Name: "Saw Gerrera",
	},
	"SOR154": {
		ID: "SOR154",
		Name: "Rallying Cry",
	},
	"SOR155": {
		ID: "SOR155",
		Name: "Aggression",
	},
	"SOR156": {
		ID: "SOR156",
		Name: "Benthic "Two Tubes"",
	},
	"SOR157": {
		ID: "SOR157",
		Name: "Cantina Braggart",
	},
	"SOR158": {
		ID: "SOR158",
		Name: "Jedha Agitator",
	},
	"SOR159": {
		ID: "SOR159",
		Name: "Partisan Insurgent",
	},
	"SOR160": {
		ID: "SOR160",
		Name: "Wolffe",
	},
	"SOR161": {
		ID: "SOR161",
		Name: "Ardent Sympathizer",
	},
	"SOR162": {
		ID: "SOR162",
		Name: "Disabling Fang Fighter",
	},
	"SOR163": {
		ID: "SOR163",
		Name: "Star Wing Scout",
	},
	"SOR164": {
		ID: "SOR164",
		Name: "Wampa",
	},
	"SOR165": {
		ID: "SOR165",
		Name: "Occupier Siege Tank",
	},
	"SOR166": {
		ID: "SOR166",
		Name: "Infiltrator's Skill",
	},
	"SOR167": {
		ID: "SOR167",
		Name: "Force Throw",
	},
	"SOR168": {
		ID: "SOR168",
		Name: "Precision Fire",
	},
	"SOR169": {
		ID: "SOR169",
		Name: "Keep Fighting",
	},
	"SOR170": {
		ID: "SOR170",
		Name: "Power Failure",
	},
	"SOR171": {
		ID: "SOR171",
		Name: "Mission Briefing",
	},
	"SOR172": {
		ID: "SOR172",
		Name: "Open Fire",
	},
	"SOR173": {
		ID: "SOR173",
		Name: "Bombing Run",
	},
	"SOR174": {
		ID: "SOR174",
		Name: "Smoke and Cinders",
	},
	"SOR175": {
		ID: "SOR175",
		Name: "Forced Surrender",
	},
	"SOR176": {
		ID: "SOR176",
		Name: "ISB Agent",
	},
	"SOR177": {
		ID: "SOR177",
		Name: "Bib Fortuna",
	},
	"SOR178": {
		ID: "SOR178",
		Name: "Cartel Spacer",
	},
	"SOR179": {
		ID: "SOR179",
		Name: "Boba Fett",
	},
	"SOR180": {
		ID: "SOR180",
		Name: "Seventh Fleet Defender",
	},
	"SOR181": {
		ID: "SOR181",
		Name: "Jabba the Hutt",
	},
	"SOR182": {
		ID: "SOR182",
		Name: "Bossk",
	},
	"SOR183": {
		ID: "SOR183",
		Name: "Bounty Hunter Crew",
	},
	"SOR184": {
		ID: "SOR184",
		Name: "Fett's Firespray",
	},
	"SOR185": {
		ID: "SOR185",
		Name: "Chimaera",
	},
	"SOR186": {
		ID: "SOR186",
		Name: "No Good to Me Dead",
	},
	"SOR187": {
		ID: "SOR187",
		Name: "I Had No Choice",
	},
	"SOR188": {
		ID: "SOR188",
		Name: "Chopper",
	},
	"SOR189": {
		ID: "SOR189",
		Name: "Leia Organa",
	},
	"SOR190": {
		ID: "SOR190",
		Name: "Lothal Insurgent",
	},
	"SOR191": {
		ID: "SOR191",
		Name: "Vanguard Ace",
	},
	"SOR192": {
		ID: "SOR192",
		Name: "Ezra Bridger",
	},
	"SOR193": {
		ID: "SOR193",
		Name: "Millennium Falcon",
	},
	"SOR194": {
		ID: "SOR194",
		Name: "Rogue Operative",
	},
	"SOR195": {
		ID: "SOR195",
		Name: "Auzituck Liberator Gunship",
	},
	"SOR196": {
		ID: "SOR196",
		Name: "Chewbacca",
	},
	"SOR197": {
		ID: "SOR197",
		Name: "Lando Calrissian",
	},
	"SOR198": {
		ID: "SOR198",
		Name: "Han Solo",
	},
	"SOR199": {
		ID: "SOR199",
		Name: "Bamboozle",
	},
	"SOR200": {
		ID: "SOR200",
		Name: "Spark of Rebellion",
	},
	"SOR201": {
		ID: "SOR201",
		Name: "Bodhi Rook",
	},
	"SOR202": {
		ID: "SOR202",
		Name: "Cantina Bouncer",
	},
	"SOR203": {
		ID: "SOR203",
		Name: "Cunning",
	},
	"SOR204": {
		ID: "SOR204",
		Name: "Greedo",
	},
	"SOR205": {
		ID: "SOR205",
		Name: "Jawa Scavenger",
	},
	"SOR206": {
		ID: "SOR206",
		Name: "Mining Guild TIE Fighter",
	},
	"SOR207": {
		ID: "SOR207",
		Name: "Crafty Smuggler",
	},
	"SOR208": {
		ID: "SOR208",
		Name: "Outer Rim Headhunter",
	},
	"SOR209": {
		ID: "SOR209",
		Name: "Pirated Starfighter",
	},
	"SOR210": {
		ID: "SOR210",
		Name: "Swoop Racer",
	},
	"SOR211": {
		ID: "SOR211",
		Name: "Gamorrean Guards",
	},
	"SOR212": {
		ID: "SOR212",
		Name: "Strafing Gunship",
	},
	"SOR213": {
		ID: "SOR213",
		Name: "Syndicate Lackeys",
	},
	"SOR214": {
		ID: "SOR214",
		Name: "Smuggling Compartment",
	},
	"SOR215": {
		ID: "SOR215",
		Name: "Snapshot Reflexes",
	},
	"SOR216": {
		ID: "SOR216",
		Name: "Disarm",
	},
	"SOR217": {
		ID: "SOR217",
		Name: "Shoot First",
	},
	"SOR218": {
		ID: "SOR218",
		Name: "Asteroid Sanctuary",
	},
	"SOR219": {
		ID: "SOR219",
		Name: "Sneak Attack",
	},
	"SOR220": {
		ID: "SOR220",
		Name: "Surprise Strike",
	},
	"SOR221": {
		ID: "SOR221",
		Name: "Outmaneuver",
	},
	"SOR222": {
		ID: "SOR222",
		Name: "Waylay",
	},
	"SOR223": {
		ID: "SOR223",
		Name: "Don't Get Cocky",
	},
	"SOR224": {
		ID: "SOR224",
		Name: "Change of Heart",
	},
	"SOR225": {
		ID: "SOR225",
		Name: "TIE/ln Fighter",
	},
	"SOR226": {
		ID: "SOR226",
		Name: "Admiral Motti",
	},
	"SOR227": {
		ID: "SOR227",
		Name: "Snowtrooper Lieutenant",
	},
	"SOR228": {
		ID: "SOR228",
		Name: "Viper Probe Droid",
	},
	"SOR229": {
		ID: "SOR229",
		Name: "Cell Block Guard",
	},
	"SOR230": {
		ID: "SOR230",
		Name: "General Veers",
	},
	"SOR231": {
		ID: "SOR231",
		Name: "TIE Advanced",
	},
	"SOR232": {
		ID: "SOR232",
		Name: "AT-ST",
	},
	"SOR233": {
		ID: "SOR233",
		Name: "I Am Your Father",
	},
	"SOR234": {
		ID: "SOR234",
		Name: "Maximum Firepower",
	},
	"SOR235": {
		ID: "SOR235",
		Name: "Galactic Ambition",
	},
	"SOR236": {
		ID: "SOR236",
		Name: "R2-D2",
	},
	"SOR237": {
		ID: "SOR237",
		Name: "Alliance X-Wing",
	},
	"SOR238": {
		ID: "SOR238",
		Name: "C-3PO",
	},
	"SOR239": {
		ID: "SOR239",
		Name: "Rebel Pathfinder",
	},
	"SOR240": {
		ID: "SOR240",
		Name: "Fleet Lieutenant",
	},
	"SOR241": {
		ID: "SOR241",
		Name: "Wing Leader",
	},
	"SOR242": {
		ID: "SOR242",
		Name: "General Dodonna",
	},
	"SOR243": {
		ID: "SOR243",
		Name: "Regional Sympathizers",
	},
	"SOR244": {
		ID: "SOR244",
		Name: "Snowspeeder",
	},
	"SOR245": {
		ID: "SOR245",
		Name: "Medal Ceremony",
	},
	"SOR246": {
		ID: "SOR246",
		Name: "You're My Only Hope",
	},
	"SOR247": {
		ID: "SOR247",
		Name: "Underworld Thug",
	},
	"SOR248": {
		ID: "SOR248",
		Name: "Volunteer Soldier",
	},
	"SOR249": {
		ID: "SOR249",
		Name: "Frontier AT-RT",
	},
	"SOR250": {
		ID: "SOR250",
		Name: "Corellian Freighter",
	},
	"SOR251": {
		ID: "SOR251",
		Name: "Confiscate",
	},
	"SOR252": {
		ID: "SOR252",
		Name: "Restock",
	},
	"SHD001": {
		ID: "SHD001",
		Name: "Gar Saxon",
	},
	"SHD002": {
		ID: "SHD002",
		Name: "Qi'ra",
	},
	"SHD003": {
		ID: "SHD003",
		Name: "Finn",
	},
	"SHD004": {
		ID: "SHD004",
		Name: "Rey",
	},
	"SHD005": {
		ID: "SHD005",
		Name: "Hondo Ohnaka",
	},
	"SHD006": {
		ID: "SHD006",
		Name: "Jabba the Hutt",
	},
	"SHD007": {
		ID: "SHD007",
		Name: "Moff Gideon",
	},
	"SHD008": {
		ID: "SHD008",
		Name: "Boba Fett",
	},
	"SHD009": {
		ID: "SHD009",
		Name: "Hunter",
	},
	"SHD010": {
		ID: "SHD010",
		Name: "Bossk",
	},
	"SHD011": {
		ID: "SHD011",
		Name: "Kylo Ren",
	},
	"SHD012": {
		ID: "SHD012",
		Name: "Bo-Katan Kryze",
	},
	"SHD013": {
		ID: "SHD013",
		Name: "Han Solo",
	},
	"SHD014": {
		ID: "SHD014",
		Name: "Cad Bane",
	},
	"SHD015": {
		ID: "SHD015",
		Name: "Doctor Aphra",
	},
	"SHD016": {
		ID: "SHD016",
		Name: "Fennec Shand",
	},
	"SHD017": {
		ID: "SHD017",
		Name: "Lando Calrissian",
	},
	"SHD018": {
		ID: "SHD018",
		Name: "The Mandalorian",
	},
	"SHD019": {
		ID: "SHD019",
		Name: "Remnant Science Facility",
	},
	"SHD020": {
		ID: "SHD020",
		Name: "Remote Village",
	},
	"SHD021": {
		ID: "SHD021",
		Name: "Maz Kanata's Castle",
	},
	"SHD022": {
		ID: "SHD022",
		Name: "Nevarro City",
	},
	"SHD023": {
		ID: "SHD023",
		Name: "Death Watch Hideout",
	},
	"SHD024": {
		ID: "SHD024",
		Name: "Spice Mines",
	},
	"SHD025": {
		ID: "SHD025",
		Name: "Coronet City",
	},
	"SHD026": {
		ID: "SHD026",
		Name: "Jabba's Palace",
	},
	"SHD027": {
		ID: "SHD027",
		Name: "Hylobon Enforcer",
	},
	"SHD028": {
		ID: "SHD028",
		Name: "Doctor Pershing",
	},
	"SHD029": {
		ID: "SHD029",
		Name: "Pyke Sentinel",
	},
	"SHD030": {
		ID: "SHD030",
		Name: "Death Trooper",
	},
	"SHD031": {
		ID: "SHD031",
		Name: "The Client",
	},
	"SHD032": {
		ID: "SHD032",
		Name: "Lom Pyke",
	},
	"SHD033": {
		ID: "SHD033",
		Name: "Synara San",
	},
	"SHD034": {
		ID: "SHD034",
		Name: "Supercommando Squad",
	},
	"SHD035": {
		ID: "SHD035",
		Name: "Clan Saxon Gauntlet",
	},
	"SHD036": {
		ID: "SHD036",
		Name: "First Light",
	},
	"SHD037": {
		ID: "SHD037",
		Name: "Supreme Leader Snoke",
	},
	"SHD038": {
		ID: "SHD038",
		Name: "Brutal Traditions",
	},
	"SHD039": {
		ID: "SHD039",
		Name: "Calculated Lethality",
	},
	"SHD040": {
		ID: "SHD040",
		Name: "Clan Wren Rescuer",
	},
	"SHD041": {
		ID: "SHD041",
		Name: "Kuiil",
	},
	"SHD042": {
		ID: "SHD042",
		Name: "Concord Dawn Interceptors",
	},
	"SHD043": {
		ID: "SHD043",
		Name: "Village Protectors",
	},
	"SHD044": {
		ID: "SHD044",
		Name: "Razor Crest",
	},
	"SHD045": {
		ID: "SHD045",
		Name: "Rose Tico",
	},
	"SHD046": {
		ID: "SHD046",
		Name: "Rey",
	},
	"SHD047": {
		ID: "SHD047",
		Name: "The Armorer",
	},
	"SHD048": {
		ID: "SHD048",
		Name: "Gentle Giant",
	},
	"SHD049": {
		ID: "SHD049",
		Name: "The Mandalorian",
	},
	"SHD050": {
		ID: "SHD050",
		Name: "Chewbacca",
	},
	"SHD051": {
		ID: "SHD051",
		Name: "Mystic Reflection",
	},
	"SHD052": {
		ID: "SHD052",
		Name: "Sugi",
	},
	"SHD053": {
		ID: "SHD053",
		Name: "Second Chance",
	},
	"SHD054": {
		ID: "SHD054",
		Name: "Midnight Repairs",
	},
	"SHD055": {
		ID: "SHD055",
		Name: "Moisture Farmer",
	},
	"SHD056": {
		ID: "SHD056",
		Name: "Follower of The Way",
	},
	"SHD057": {
		ID: "SHD057",
		Name: "Rickety Quadjumper",
	},
	"SHD058": {
		ID: "SHD058",
		Name: "Val",
	},
	"SHD059": {
		ID: "SHD059",
		Name: "Embo",
	},
	"SHD060": {
		ID: "SHD060",
		Name: "HWK-290 Freighter",
	},
	"SHD061": {
		ID: "SHD061",
		Name: "Wroshyr Tree Tender",
	},
	"SHD062": {
		ID: "SHD062",
		Name: "Niima Outpost Constables",
	},
	"SHD063": {
		ID: "SHD063",
		Name: "System Patrol Craft",
	},
	"SHD064": {
		ID: "SHD064",
		Name: "Survivors' Gauntlet",
	},
	"SHD065": {
		ID: "SHD065",
		Name: "Vigilant Pursuit Craft",
	},
	"SHD066": {
		ID: "SHD066",
		Name: "Cargo Juggernaut",
	},
	"SHD067": {
		ID: "SHD067",
		Name: "Fenn Rau",
	},
	"SHD068": {
		ID: "SHD068",
		Name: "Public Enemy",
	},
	"SHD069": {
		ID: "SHD069",
		Name: "Foundling",
	},
	"SHD070": {
		ID: "SHD070",
		Name: "Resilient",
	},
	"SHD071": {
		ID: "SHD071",
		Name: "Top Target",
	},
	"SHD072": {
		ID: "SHD072",
		Name: "Imprisoned",
	},
	"SHD073": {
		ID: "SHD073",
		Name: "Mandalorian Armor",
	},
	"SHD074": {
		ID: "SHD074",
		Name: "Vambrace Grappleshot",
	},
	"SHD075": {
		ID: "SHD075",
		Name: "Covert Strength",
	},
	"SHD076": {
		ID: "SHD076",
		Name: "Unexpected Escape",
	},
	"SHD077": {
		ID: "SHD077",
		Name: "Evidence of the Crime",
	},
	"SHD078": {
		ID: "SHD078",
		Name: "Fell the Dragon",
	},
	"SHD079": {
		ID: "SHD079",
		Name: "Rival's Fall",
	},
	"SHD080": {
		ID: "SHD080",
		Name: "Salacious Crumb",
	},
	"SHD081": {
		ID: "SHD081",
		Name: "General Tagge",
	},
	"SHD082": {
		ID: "SHD082",
		Name: "Outland TIE Vanguard",
	},
	"SHD083": {
		ID: "SHD083",
		Name: "Seasoned Shoretrooper",
	},
	"SHD084": {
		ID: "SHD084",
		Name: "Phase-III Dark Trooper",
	},
	"SHD085": {
		ID: "SHD085",
		Name: "Superlaser Technician",
	},
	"SHD086": {
		ID: "SHD086",
		Name: "Warbird Stowaway",
	},
	"SHD087": {
		ID: "SHD087",
		Name: "Crosshair",
	},
	"SHD088": {
		ID: "SHD088",
		Name: "Ephant Mon",
	},
	"SHD089": {
		ID: "SHD089",
		Name: "Pirate Battle Tank",
	},
	"SHD090": {
		ID: "SHD090",
		Name: "Maul",
	},
	"SHD091": {
		ID: "SHD091",
		Name: "Jabba's Rancor",
	},
	"SHD092": {
		ID: "SHD092",
		Name: "Finalizer",
	},
	"SHD093": {
		ID: "SHD093",
		Name: "Remnant Reserves",
	},
	"SHD094": {
		ID: "SHD094",
		Name: "Palpatine's Return",
	},
	"SHD095": {
		ID: "SHD095",
		Name: "Clone Deserter",
	},
	"SHD096": {
		ID: "SHD096",
		Name: "Maz Kanata",
	},
	"SHD097": {
		ID: "SHD097",
		Name: "Freetown Backup",
	},
	"SHD098": {
		ID: "SHD098",
		Name: "Sundari Peacekeeper",
	},
	"SHD099": {
		ID: "SHD099",
		Name: "Echo",
	},
	"SHD100": {
		ID: "SHD100",
		Name: "Modded Cohort",
	},
	"SHD101": {
		ID: "SHD101",
		Name: "Adelphi Patrol Wing",
	},
	"SHD102": {
		ID: "SHD102",
		Name: "The Marauder",
	},
	"SHD103": {
		ID: "SHD103",
		Name: "General Rieekan",
	},
	"SHD104": {
		ID: "SHD104",
		Name: "Inspiring Mentor",
	},
	"SHD105": {
		ID: "SHD105",
		Name: "Spark of Hope",
	},
	"SHD106": {
		ID: "SHD106",
		Name: "Rule with Respect",
	},
	"SHD107": {
		ID: "SHD107",
		Name: "Enterprising Lackeys",
	},
	"SHD108": {
		ID: "SHD108",
		Name: "Enforced Loyalty",
	},
	"SHD109": {
		ID: "SHD109",
		Name: "Endless Legions",
	},
	"SHD110": {
		ID: "SHD110",
		Name: "Warzone Lieutenant",
	},
	"SHD111": {
		ID: "SHD111",
		Name: "Collections Starhopper",
	},
	"SHD112": {
		ID: "SHD112",
		Name: "Gamorrean Retainer",
	},
	"SHD113": {
		ID: "SHD113",
		Name: "Privateer Crew",
	},
	"SHD114": {
		ID: "SHD114",
		Name: "Scanning Officer",
	},
	"SHD115": {
		ID: "SHD115",
		Name: "Cobb Vanth",
	},
	"SHD116": {
		ID: "SHD116",
		Name: "Outlaw Corona",
	},
	"SHD117": {
		ID: "SHD117",
		Name: "Reputable Hunter",
	},
	"SHD118": {
		ID: "SHD118",
		Name: "Kihraxz Heavy Fighter",
	},
	"SHD119": {
		ID: "SHD119",
		Name: "Weequay Pirate Gang",
	},
	"SHD120": {
		ID: "SHD120",
		Name: "Discerning Veteran",
	},
	"SHD121": {
		ID: "SHD121",
		Name: "Mercenary Company",
	},
	"SHD122": {
		ID: "SHD122",
		Name: "Arquitens Assault Cruiser",
	},
	"SHD123": {
		ID: "SHD123",
		Name: "Bounty Hunter's Quarry",
	},
	"SHD124": {
		ID: "SHD124",
		Name: "Legal Authority",
	},
	"SHD125": {
		ID: "SHD125",
		Name: "Price on Your Head",
	},
	"SHD126": {
		ID: "SHD126",
		Name: "The Darksaber",
	},
	"SHD127": {
		ID: "SHD127",
		Name: "Commission",
	},
	"SHD128": {
		ID: "SHD128",
		Name: "Outflank",
	},
	"SHD129": {
		ID: "SHD129",
		Name: "Timely Intervention",
	},
	"SHD130": {
		ID: "SHD130",
		Name: "Moment of Glory",
	},
	"SHD131": {
		ID: "SHD131",
		Name: "Take Captive",
	},
	"SHD132": {
		ID: "SHD132",
		Name: "Choose Sides",
	},
	"SHD133": {
		ID: "SHD133",
		Name: "Dengar",
	},
	"SHD134": {
		ID: "SHD134",
		Name: "Guavian Antagonizer",
	},
	"SHD135": {
		ID: "SHD135",
		Name: "Kylo's TIE Silencer",
	},
	"SHD136": {
		ID: "SHD136",
		Name: "Death Watch Loyalist",
	},
	"SHD137": {
		ID: "SHD137",
		Name: "Punishing One",
	},
	"SHD138": {
		ID: "SHD138",
		Name: "Jango Fett",
	},
	"SHD139": {
		ID: "SHD139",
		Name: "Krrsantan",
	},
	"SHD140": {
		ID: "SHD140",
		Name: "Trandoshan Hunters",
	},
	"SHD141": {
		ID: "SHD141",
		Name: "Kylo Ren",
	},
	"SHD142": {
		ID: "SHD142",
		Name: "Pre Vizsla",
	},
	"SHD143": {
		ID: "SHD143",
		Name: "Ruthlessness",
	},
	"SHD144": {
		ID: "SHD144",
		Name: "Give In to Your Anger",
	},
	"SHD145": {
		ID: "SHD145",
		Name: "Headhunting",
	},
	"SHD146": {
		ID: "SHD146",
		Name: "Heroic Renegade",
	},
	"SHD147": {
		ID: "SHD147",
		Name: "Ketsu Onyo",
	},
	"SHD148": {
		ID: "SHD148",
		Name: "Cassian Andor",
	},
	"SHD149": {
		ID: "SHD149",
		Name: "Nite Owl Skirmisher",
	},
	"SHD150": {
		ID: "SHD150",
		Name: "Koska Reeves",
	},
	"SHD151": {
		ID: "SHD151",
		Name: "Valiant Assault Ship",
	},
	"SHD152": {
		ID: "SHD152",
		Name: "Desperado Freighter",
	},
	"SHD153": {
		ID: "SHD153",
		Name: "Poe Dameron",
	},
	"SHD154": {
		ID: "SHD154",
		Name: "Wrecker",
	},
	"SHD155": {
		ID: "SHD155",
		Name: "Heroic Resolve",
	},
	"SHD156": {
		ID: "SHD156",
		Name: "Cripple Authority",
	},
	"SHD157": {
		ID: "SHD157",
		Name: "Bo-Katan Kryze",
	},
	"SHD158": {
		ID: "SHD158",
		Name: "Wild Rancor",
	},
	"SHD159": {
		ID: "SHD159",
		Name: "The Chaos of War",
	},
	"SHD160": {
		ID: "SHD160",
		Name: "Reckless Gunslinger",
	},
	"SHD161": {
		ID: "SHD161",
		Name: "Stolen Landspeeder",
	},
	"SHD162": {
		ID: "SHD162",
		Name: "House Kast Soldier",
	},
	"SHD163": {
		ID: "SHD163",
		Name: "Migs Mayfeld",
	},
	"SHD164": {
		ID: "SHD164",
		Name: "Rhokai Gunship",
	},
	"SHD165": {
		ID: "SHD165",
		Name: "Unlicensed Headhunter",
	},
	"SHD166": {
		ID: "SHD166",
		Name: "Disabling Fang Fighter",
	},
	"SHD167": {
		ID: "SHD167",
		Name: "Wanted Insurgents",
	},
	"SHD168": {
		ID: "SHD168",
		Name: "Hunting Nexu",
	},
	"SHD169": {
		ID: "SHD169",
		Name: "Clan Challengers",
	},
	"SHD170": {
		ID: "SHD170",
		Name: "IG-11",
	},
	"SHD171": {
		ID: "SHD171",
		Name: "Covetous Rivals",
	},
	"SHD172": {
		ID: "SHD172",
		Name: "Krayt Dragon",
	},
	"SHD173": {
		ID: "SHD173",
		Name: "Guild Target",
	},
	"SHD174": {
		ID: "SHD174",
		Name: "Hotshot DL-44 Blaster",
	},
	"SHD175": {
		ID: "SHD175",
		Name: "Armed to the Teeth",
	},
	"SHD176": {
		ID: "SHD176",
		Name: "Death Mark",
	},
	"SHD177": {
		ID: "SHD177",
		Name: "Vambrace Flamethrower",
	},
	"SHD178": {
		ID: "SHD178",
		Name: "Daring Raid",
	},
	"SHD179": {
		ID: "SHD179",
		Name: "Desperate Attack",
	},
	"SHD180": {
		ID: "SHD180",
		Name: "Detention Block Rescue",
	},
	"SHD181": {
		ID: "SHD181",
		Name: "Pillage",
	},
	"SHD182": {
		ID: "SHD182",
		Name: "Bravado",
	},
	"SHD183": {
		ID: "SHD183",
		Name: "Kintan Intimidator",
	},
	"SHD184": {
		ID: "SHD184",
		Name: "Bazine Netal",
	},
	"SHD185": {
		ID: "SHD185",
		Name: "Doctor Evazan",
	},
	"SHD186": {
		ID: "SHD186",
		Name: "Hunter of the Haxion Brood",
	},
	"SHD187": {
		ID: "SHD187",
		Name: "Lurking TIE Phantom",
	},
	"SHD188": {
		ID: "SHD188",
		Name: "4-LOM",
	},
	"SHD189": {
		ID: "SHD189",
		Name: "Slaver's Freighter",
	},
	"SHD190": {
		ID: "SHD190",
		Name: "Zuckuss",
	},
	"SHD191": {
		ID: "SHD191",
		Name: "Xanadu Blood",
	},
	"SHD192": {
		ID: "SHD192",
		Name: "Dryden Vos",
	},
	"SHD193": {
		ID: "SHD193",
		Name: "Frozen in Carbonite",
	},
	"SHD194": {
		ID: "SHD194",
		Name: "Triple Dark Raid",
	},
	"SHD195": {
		ID: "SHD195",
		Name: "Cartel Turncoat",
	},
	"SHD196": {
		ID: "SHD196",
		Name: "Grogu",
	},
	"SHD197": {
		ID: "SHD197",
		Name: "L3-37",
	},
	"SHD198": {
		ID: "SHD198",
		Name: "Omega",
	},
	"SHD199": {
		ID: "SHD199",
		Name: "Coruscant Dissident",
	},
	"SHD200": {
		ID: "SHD200",
		Name: "Liberated Slaves",
	},
	"SHD201": {
		ID: "SHD201",
		Name: "Principled Outlaw",
	},
	"SHD202": {
		ID: "SHD202",
		Name: "Qi'ra",
	},
	"SHD203": {
		ID: "SHD203",
		Name: "Zorii Bliss",
	},
	"SHD204": {
		ID: "SHD204",
		Name: "Millennium Falcon",
	},
	"SHD205": {
		ID: "SHD205",
		Name: "Let the Wookiee Win",
	},
	"SHD206": {
		ID: "SHD206",
		Name: "Spare the Target",
	},
	"SHD207": {
		ID: "SHD207",
		Name: "A New Adventure",
	},
	"SHD208": {
		ID: "SHD208",
		Name: "Final Showdown",
	},
	"SHD209": {
		ID: "SHD209",
		Name: "Criminal Muscle",
	},
	"SHD210": {
		ID: "SHD210",
		Name: "Cloud-Rider",
	},
	"SHD211": {
		ID: "SHD211",
		Name: "Fugitive Wookiee",
	},
	"SHD212": {
		ID: "SHD212",
		Name: "Privateer Scyk",
	},
	"SHD213": {
		ID: "SHD213",
		Name: "DJ",
	},
	"SHD214": {
		ID: "SHD214",
		Name: "Frontier Trader",
	},
	"SHD215": {
		ID: "SHD215",
		Name: "Smuggler's Starfighter",
	},
	"SHD216": {
		ID: "SHD216",
		Name: "Chain Code Collector",
	},
	"SHD217": {
		ID: "SHD217",
		Name: "Tobias Beckett",
	},
	"SHD218": {
		ID: "SHD218",
		Name: "Resourceful Pursuers",
	},
	"SHD219": {
		ID: "SHD219",
		Name: "Enfys Nest",
	},
	"SHD220": {
		ID: "SHD220",
		Name: "Fennec Shand",
	},
	"SHD221": {
		ID: "SHD221",
		Name: "Wanted",
	},
	"SHD222": {
		ID: "SHD222",
		Name: "Enticing Reward",
	},
	"SHD223": {
		ID: "SHD223",
		Name: "Snapshot Reflexes",
	},
	"SHD224": {
		ID: "SHD224",
		Name: "Boba Fett's Armor",
	},
	"SHD225": {
		ID: "SHD225",
		Name: "Jetpack",
	},
	"SHD226": {
		ID: "SHD226",
		Name: "Unrefusable Offer",
	},
	"SHD227": {
		ID: "SHD227",
		Name: "Look the Other Way",
	},
	"SHD228": {
		ID: "SHD228",
		Name: "Bounty Posting",
	},
	"SHD229": {
		ID: "SHD229",
		Name: "Ma Klounkee",
	},
	"SHD230": {
		ID: "SHD230",
		Name: "Swoop Down",
	},
	"SHD231": {
		ID: "SHD231",
		Name: "Surprise Strike",
	},
	"SHD232": {
		ID: "SHD232",
		Name: "Relentless Pursuit",
	},
	"SHD233": {
		ID: "SHD233",
		Name: "Evacuate",
	},
	"SHD234": {
		ID: "SHD234",
		Name: "Incinerator Trooper",
	},
	"SHD235": {
		ID: "SHD235",
		Name: "Ruthless Assassin",
	},
	"SHD236": {
		ID: "SHD236",
		Name: "Snowtrooper Lieutenant",
	},
	"SHD237": {
		ID: "SHD237",
		Name: "Black Sun Starfighter",
	},
	"SHD238": {
		ID: "SHD238",
		Name: "Cell Block Guard",
	},
	"SHD239": {
		ID: "SHD239",
		Name: "Toro Calican",
	},
	"SHD240": {
		ID: "SHD240",
		Name: "Hutt's Henchmen",
	},
	"SHD241": {
		ID: "SHD241",
		Name: "Kragan Gorr",
	},
	"SHD242": {
		ID: "SHD242",
		Name: "Gideon's Light Cruiser",
	},
	"SHD243": {
		ID: "SHD243",
		Name: "Altering the Deal",
	},
	"SHD244": {
		ID: "SHD244",
		Name: "No Bargain",
	},
	"SHD245": {
		ID: "SHD245",
		Name: "Greef Karga",
	},
	"SHD246": {
		ID: "SHD246",
		Name: "Grey Squadron Y-Wing",
	},
	"SHD247": {
		ID: "SHD247",
		Name: "Protector of the Throne",
	},
	"SHD248": {
		ID: "SHD248",
		Name: "Tech",
	},
	"SHD249": {
		ID: "SHD249",
		Name: "Wookiee Warrior",
	},
	"SHD250": {
		ID: "SHD250",
		Name: "Tarfful",
	},
	"SHD251": {
		ID: "SHD251",
		Name: "The Mandalorian's Rifle",
	},
	"SHD252": {
		ID: "SHD252",
		Name: "Smuggler's Aid",
	},
	"SHD253": {
		ID: "SHD253",
		Name: "This Is The Way",
	},
	"SHD254": {
		ID: "SHD254",
		Name: "Bounty Guild Initiate",
	},
	"SHD255": {
		ID: "SHD255",
		Name: "Lady Proxima",
	},
	"SHD256": {
		ID: "SHD256",
		Name: "Mercenary Gunship",
	},
	"SHD257": {
		ID: "SHD257",
		Name: "Underworld Thug",
	},
	"SHD258": {
		ID: "SHD258",
		Name: "Mandalorian Warrior",
	},
	"SHD259": {
		ID: "SHD259",
		Name: "Twin Pod Cloud Car",
	},
	"SHD260": {
		ID: "SHD260",
		Name: "Street Gang Recruiter",
	},
	"SHD261": {
		ID: "SHD261",
		Name: "Rich Reward",
	},
	"SHD262": {
		ID: "SHD262",
		Name: "Confiscate",
	},
	"TWI001": {
		ID: "TWI001",
		Name: "Nala Se",
	},
	"TWI002": {
		ID: "TWI002",
		Name: "Nute Gunray",
	},
	"TWI003": {
		ID: "TWI003",
		Name: "Obi-Wan Kenobi",
	},
	"TWI004": {
		ID: "TWI004",
		Name: "Yoda",
	},
	"TWI005": {
		ID: "TWI005",
		Name: "Count Dooku",
	},
	"TWI006": {
		ID: "TWI006",
		Name: "Wat Tambor",
	},
	"TWI007": {
		ID: "TWI007",
		Name: "Captain Rex",
	},
	"TWI008": {
		ID: "TWI008",
		Name: "Padmé Amidala",
	},
	"TWI009": {
		ID: "TWI009",
		Name: "Maul",
	},
	"TWI010": {
		ID: "TWI010",
		Name: "Pre Vizsla",
	},
	"TWI011": {
		ID: "TWI011",
		Name: "Ahsoka Tano",
	},
	"TWI012": {
		ID: "TWI012",
		Name: "Anakin Skywalker",
	},
	"TWI013": {
		ID: "TWI013",
		Name: "Mace Windu",
	},
	"TWI014": {
		ID: "TWI014",
		Name: "Asajj Ventress",
	},
	"TWI015": {
		ID: "TWI015",
		Name: "General Grievous",
	},
	"TWI016": {
		ID: "TWI016",
		Name: "Jango Fett",
	},
	"TWI017": {
		ID: "TWI017",
		Name: "Chancellor Palpatine",
	},
	"TWI018": {
		ID: "TWI018",
		Name: "Quinlan Vos",
	},
	"TWI019": {
		ID: "TWI019",
		Name: "Pau City",
	},
	"TWI020": {
		ID: "TWI020",
		Name: "Sundari",
	},
	"TWI021": {
		ID: "TWI021",
		Name: "The Crystal City",
	},
	"TWI022": {
		ID: "TWI022",
		Name: "Droid Manufactory",
	},
	"TWI023": {
		ID: "TWI023",
		Name: "Lair of Grievous",
	},
	"TWI024": {
		ID: "TWI024",
		Name: "Tipoca City",
	},
	"TWI025": {
		ID: "TWI025",
		Name: "Shadow Collective Camp",
	},
	"TWI026": {
		ID: "TWI026",
		Name: "KCM Mining Facility",
	},
	"TWI027": {
		ID: "TWI027",
		Name: "The Nest",
	},
	"TWI028": {
		ID: "TWI028",
		Name: "Petranaki Arena",
	},
	"TWI029": {
		ID: "TWI029",
		Name: "Level 1313",
	},
	"TWI030": {
		ID: "TWI030",
		Name: "Pyke Palace",
	},
	"TWI031": {
		ID: "TWI031",
		Name: "Rune Haako",
	},
	"TWI032": {
		ID: "TWI032",
		Name: "Wartime Trade Official",
	},
	"TWI033": {
		ID: "TWI033",
		Name: "Calculating MagnaGuard",
	},
	"TWI034": {
		ID: "TWI034",
		Name: "General Grievous",
	},
	"TWI035": {
		ID: "TWI035",
		Name: "Morgan Elsbeth",
	},
	"TWI036": {
		ID: "TWI036",
		Name: "Devastating Gunship",
	},
	"TWI037": {
		ID: "TWI037",
		Name: "Droideka Security",
	},
	"TWI038": {
		ID: "TWI038",
		Name: "Providence Destroyer",
	},
	"TWI039": {
		ID: "TWI039",
		Name: "Malevolence",
	},
	"TWI040": {
		ID: "TWI040",
		Name: "A Fine Addition",
	},
	"TWI041": {
		ID: "TWI041",
		Name: "Lethal Crackdown",
	},
	"TWI042": {
		ID: "TWI042",
		Name: "Barriss Offee",
	},
	"TWI043": {
		ID: "TWI043",
		Name: "Outspoken Representative",
	},
	"TWI044": {
		ID: "TWI044",
		Name: "Kashyyyk Defender",
	},
	"TWI045": {
		ID: "TWI045",
		Name: "41st Elite Corps",
	},
	"TWI046": {
		ID: "TWI046",
		Name: "Captain Typho",
	},
	"TWI047": {
		ID: "TWI047",
		Name: "Satine Kryze",
	},
	"TWI048": {
		ID: "TWI048",
		Name: "Obi-Wan's Aethersprite",
	},
	"TWI049": {
		ID: "TWI049",
		Name: "Knight of the Republic",
	},
	"TWI050": {
		ID: "TWI050",
		Name: "Luminara Unduli",
	},
	"TWI051": {
		ID: "TWI051",
		Name: "For the Republic",
	},
	"TWI052": {
		ID: "TWI052",
		Name: "Hello There",
	},
	"TWI053": {
		ID: "TWI053",
		Name: "Finn",
	},
	"TWI054": {
		ID: "TWI054",
		Name: "Duchess's Champion",
	},
	"TWI055": {
		ID: "TWI055",
		Name: "Equalize",
	},
	"TWI056": {
		ID: "TWI056",
		Name: "Compassionate Senator",
	},
	"TWI057": {
		ID: "TWI057",
		Name: "Warrior Drone",
	},
	"TWI058": {
		ID: "TWI058",
		Name: "Padawan Starfighter",
	},
	"TWI059": {
		ID: "TWI059",
		Name: "Royal Guard Attaché",
	},
	"TWI060": {
		ID: "TWI060",
		Name: "Trade Federation Shuttle",
	},
	"TWI061": {
		ID: "TWI061",
		Name: "Infantry of the 212th",
	},
	"TWI062": {
		ID: "TWI062",
		Name: "Daughter of Dathomir",
	},
	"TWI063": {
		ID: "TWI063",
		Name: "Vulture Interceptor Wing",
	},
	"TWI064": {
		ID: "TWI064",
		Name: "Ki-Adi-Mundi",
	},
	"TWI065": {
		ID: "TWI065",
		Name: "Falchion Ion Tank",
	},
	"TWI066": {
		ID: "TWI066",
		Name: "Multi-Troop Transport",
	},
	"TWI067": {
		ID: "TWI067",
		Name: "The Zillo Beast",
	},
	"TWI068": {
		ID: "TWI068",
		Name: "Foresight",
	},
	"TWI069": {
		ID: "TWI069",
		Name: "Roger Roger",
	},
	"TWI070": {
		ID: "TWI070",
		Name: "Perilous Position",
	},
	"TWI071": {
		ID: "TWI071",
		Name: "Unshakeable Will",
	},
	"TWI072": {
		ID: "TWI072",
		Name: "I Have the High Ground",
	},
	"TWI073": {
		ID: "TWI073",
		Name: "Grievous Reassembly",
	},
	"TWI074": {
		ID: "TWI074",
		Name: "Guarding the Way",
	},
	"TWI075": {
		ID: "TWI075",
		Name: "Disruptive Burst",
	},
	"TWI076": {
		ID: "TWI076",
		Name: "Death by Droids",
	},
	"TWI077": {
		ID: "TWI077",
		Name: "Vanquish",
	},
	"TWI078": {
		ID: "TWI078",
		Name: "The Invasion of Christophsis",
	},
	"TWI079": {
		ID: "TWI079",
		Name: "Confederate Courier",
	},
	"TWI080": {
		ID: "TWI080",
		Name: "Poggle the Lesser",
	},
	"TWI081": {
		ID: "TWI081",
		Name: "Droid Commando",
	},
	"TWI082": {
		ID: "TWI082",
		Name: "MagnaGuard Wing Leader",
	},
	"TWI083": {
		ID: "TWI083",
		Name: "General's Guardian",
	},
	"TWI084": {
		ID: "TWI084",
		Name: "Kraken",
	},
	"TWI085": {
		ID: "TWI085",
		Name: "Kalani",
	},
	"TWI086": {
		ID: "TWI086",
		Name: "Admiral Trench",
	},
	"TWI087": {
		ID: "TWI087",
		Name: "Separatist Super Tank",
	},
	"TWI088": {
		ID: "TWI088",
		Name: "Reprocess",
	},
	"TWI089": {
		ID: "TWI089",
		Name: "Consolidation of Power",
	},
	"TWI090": {
		ID: "TWI090",
		Name: "Echo",
	},
	"TWI091": {
		ID: "TWI091",
		Name: "Republic Tactical Officer",
	},
	"TWI092": {
		ID: "TWI092",
		Name: "Admiral Yularen",
	},
	"TWI093": {
		ID: "TWI093",
		Name: "Advanced Recon Commando",
	},
	"TWI094": {
		ID: "TWI094",
		Name: "Shaak Ti",
	},
	"TWI095": {
		ID: "TWI095",
		Name: "Pelta Supply Frigate",
	},
	"TWI096": {
		ID: "TWI096",
		Name: "Aayla Secura",
	},
	"TWI097": {
		ID: "TWI097",
		Name: "Captain Rex",
	},
	"TWI098": {
		ID: "TWI098",
		Name: "Republic Defense Carrier",
	},
	"TWI099": {
		ID: "TWI099",
		Name: "Synchronized Strike",
	},
	"TWI100": {
		ID: "TWI100",
		Name: "Petition the Senate",
	},
	"TWI101": {
		ID: "TWI101",
		Name: "Mas Amedda",
	},
	"TWI102": {
		ID: "TWI102",
		Name: "Manufactured Soldiers",
	},
	"TWI103": {
		ID: "TWI103",
		Name: "Pyrrhic Assault",
	},
	"TWI104": {
		ID: "TWI104",
		Name: "Obedient Vanguard",
	},
	"TWI105": {
		ID: "TWI105",
		Name: "Steadfast Senator",
	},
	"TWI106": {
		ID: "TWI106",
		Name: "Coruscant Guard",
	},
	"TWI107": {
		ID: "TWI107",
		Name: "Patrolling V-Wing",
	},
	"TWI108": {
		ID: "TWI108",
		Name: "Ryloth Militia",
	},
	"TWI109": {
		ID: "TWI109",
		Name: "501st Liberator",
	},
	"TWI110": {
		ID: "TWI110",
		Name: "Huyang",
	},
	"TWI111": {
		ID: "TWI111",
		Name: "Republic ARC-170",
	},
	"TWI112": {
		ID: "TWI112",
		Name: "Subjugating Starfighter",
	},
	"TWI113": {
		ID: "TWI113",
		Name: "B2 Legionnaires",
	},
	"TWI114": {
		ID: "TWI114",
		Name: "Clone Commander Cody",
	},
	"TWI115": {
		ID: "TWI115",
		Name: "Osi Sobeck",
	},
	"TWI116": {
		ID: "TWI116",
		Name: "Clone",
	},
	"TWI117": {
		ID: "TWI117",
		Name: "Baktoid Spider Droid",
	},
	"TWI118": {
		ID: "TWI118",
		Name: "Gor",
	},
	"TWI119": {
		ID: "TWI119",
		Name: "Nameless Valor",
	},
	"TWI120": {
		ID: "TWI120",
		Name: "Strategic Acumen",
	},
	"TWI121": {
		ID: "TWI121",
		Name: "General's Blade",
	},
	"TWI122": {
		ID: "TWI122",
		Name: "Squad Support",
	},
	"TWI123": {
		ID: "TWI123",
		Name: "Outflank",
	},
	"TWI124": {
		ID: "TWI124",
		Name: "Tactical Advantage",
	},
	"TWI125": {
		ID: "TWI125",
		Name: "The Clone Wars",
	},
	"TWI126": {
		ID: "TWI126",
		Name: "Encouraging Leadership",
	},
	"TWI127": {
		ID: "TWI127",
		Name: "Resupply",
	},
	"TWI128": {
		ID: "TWI128",
		Name: "Take Captive",
	},
	"TWI129": {
		ID: "TWI129",
		Name: "In Defense of Kamino",
	},
	"TWI130": {
		ID: "TWI130",
		Name: "Bo-Katan Kryze",
	},
	"TWI131": {
		ID: "TWI131",
		Name: "OOM-Series Officer",
	},
	"TWI132": {
		ID: "TWI132",
		Name: "Confederate Tri-Fighter",
	},
	"TWI133": {
		ID: "TWI133",
		Name: "B1 Attack Platform",
	},
	"TWI134": {
		ID: "TWI134",
		Name: "Asajj Ventress",
	},
	"TWI135": {
		ID: "TWI135",
		Name: "Darth Maul",
	},
	"TWI136": {
		ID: "TWI136",
		Name: "Squadron of Vultures",
	},
	"TWI137": {
		ID: "TWI137",
		Name: "Savage Opress",
	},
	"TWI138": {
		ID: "TWI138",
		Name: "Count Dooku",
	},
	"TWI139": {
		ID: "TWI139",
		Name: "Corner the Prey",
	},
	"TWI140": {
		ID: "TWI140",
		Name: "Self-Destruct",
	},
	"TWI141": {
		ID: "TWI141",
		Name: "Soldier of the 501st",
	},
	"TWI142": {
		ID: "TWI142",
		Name: "Anakin's Interceptor",
	},
	"TWI143": {
		ID: "TWI143",
		Name: "Jyn Erso",
	},
	"TWI144": {
		ID: "TWI144",
		Name: "Batch Brothers",
	},
	"TWI145": {
		ID: "TWI145",
		Name: "Jesse",
	},
	"TWI146": {
		ID: "TWI146",
		Name: "Steela Gerrera",
	},
	"TWI147": {
		ID: "TWI147",
		Name: "Anakin Skywalker",
	},
	"TWI148": {
		ID: "TWI148",
		Name: "Senatorial Corvette",
	},
	"TWI149": {
		ID: "TWI149",
		Name: "Low Altitude Gunship",
	},
	"TWI150": {
		ID: "TWI150",
		Name: "Saw Gerrera",
	},
	"TWI151": {
		ID: "TWI151",
		Name: "Resolute",
	},
	"TWI152": {
		ID: "TWI152",
		Name: "Mace Windu's Lightsaber",
	},
	"TWI153": {
		ID: "TWI153",
		Name: "Bold Resistance",
	},
	"TWI154": {
		ID: "TWI154",
		Name: "Mister Bones",
	},
	"TWI155": {
		ID: "TWI155",
		Name: "Twice the Pride",
	},
	"TWI156": {
		ID: "TWI156",
		Name: "Unlimited Power",
	},
	"TWI157": {
		ID: "TWI157",
		Name: "Disaffected Senator",
	},
	"TWI158": {
		ID: "TWI158",
		Name: "Clone Heavy Gunner",
	},
	"TWI159": {
		ID: "TWI159",
		Name: "Dendup's Loyalist",
	},
	"TWI160": {
		ID: "TWI160",
		Name: "Vanguard Droid Bomber",
	},
	"TWI161": {
		ID: "TWI161",
		Name: "Bold Recon Commando",
	},
	"TWI162": {
		ID: "TWI162",
		Name: "Reckless Torrent",
	},
	"TWI163": {
		ID: "TWI163",
		Name: "Relentless Rocket Droid",
	},
	"TWI164": {
		ID: "TWI164",
		Name: "Hevy",
	},
	"TWI165": {
		ID: "TWI165",
		Name: "Kit Fisto",
	},
	"TWI166": {
		ID: "TWI166",
		Name: "Aurra Sing",
	},
	"TWI167": {
		ID: "TWI167",
		Name: "Heavy Persuader Tank",
	},
	"TWI168": {
		ID: "TWI168",
		Name: "Old Access Codes",
	},
	"TWI169": {
		ID: "TWI169",
		Name: "Clone Cohort",
	},
	"TWI170": {
		ID: "TWI170",
		Name: "Daring Raid",
	},
	"TWI171": {
		ID: "TWI171",
		Name: "Grenade Strike",
	},
	"TWI172": {
		ID: "TWI172",
		Name: "Grim Resolve",
	},
	"TWI173": {
		ID: "TWI173",
		Name: "Blood Sport",
	},
	"TWI174": {
		ID: "TWI174",
		Name: "Open Fire",
	},
	"TWI175": {
		ID: "TWI175",
		Name: "Strategic Analysis",
	},
	"TWI176": {
		ID: "TWI176",
		Name: "Caught in the Crossfire",
	},
	"TWI177": {
		ID: "TWI177",
		Name: "Guerilla Insurgency",
	},
	"TWI178": {
		ID: "TWI178",
		Name: "Planetary Invasion",
	},
	"TWI179": {
		ID: "TWI179",
		Name: "Soulless One",
	},
	"TWI180": {
		ID: "TWI180",
		Name: "Separatist Commando",
	},
	"TWI181": {
		ID: "TWI181",
		Name: "Elite P-38 Starfighter",
	},
	"TWI182": {
		ID: "TWI182",
		Name: "Infiltrating Demolisher",
	},
	"TWI183": {
		ID: "TWI183",
		Name: "Rush Clovis",
	},
	"TWI184": {
		ID: "TWI184",
		Name: "Tactical Droid Commander",
	},
	"TWI185": {
		ID: "TWI185",
		Name: "Ziro the Hutt",
	},
	"TWI186": {
		ID: "TWI186",
		Name: "San Hill",
	},
	"TWI187": {
		ID: "TWI187",
		Name: "Cad Bane",
	},
	"TWI188": {
		ID: "TWI188",
		Name: "Wartime Profiteering",
	},
	"TWI189": {
		ID: "TWI189",
		Name: "Unnatural Life",
	},
	"TWI190": {
		ID: "TWI190",
		Name: "On the Doorstep",
	},
	"TWI191": {
		ID: "TWI191",
		Name: "Wolf Pack Escort",
	},
	"TWI192": {
		ID: "TWI192",
		Name: "Padmé Amidala",
	},
	"TWI193": {
		ID: "TWI193",
		Name: "R2-D2",
	},
	"TWI194": {
		ID: "TWI194",
		Name: "Ahsoka Tano",
	},
	"TWI195": {
		ID: "TWI195",
		Name: "Sabine Wren",
	},
	"TWI196": {
		ID: "TWI196",
		Name: "Plo Koon",
	},
	"TWI197": {
		ID: "TWI197",
		Name: "Republic Attack Pod",
	},
	"TWI198": {
		ID: "TWI198",
		Name: "Enfys Nest",
	},
	"TWI199": {
		ID: "TWI199",
		Name: "Clear the Field",
	},
	"TWI200": {
		ID: "TWI200",
		Name: "Creative Thinking",
	},
	"TWI201": {
		ID: "TWI201",
		Name: "Aid from the Innocent",
	},
	"TWI202": {
		ID: "TWI202",
		Name: "Jar Jar Binks",
	},
	"TWI203": {
		ID: "TWI203",
		Name: "Chancellor Palpatine",
	},
	"TWI204": {
		ID: "TWI204",
		Name: "Impropriety Among Thieves",
	},
	"TWI205": {
		ID: "TWI205",
		Name: "Clone Dive Trooper",
	},
	"TWI206": {
		ID: "TWI206",
		Name: "Independent Senator",
	},
	"TWI207": {
		ID: "TWI207",
		Name: "B1 Security Team",
	},
	"TWI208": {
		ID: "TWI208",
		Name: "Favorable Delegate",
	},
	"TWI209": {
		ID: "TWI209",
		Name: "Hotshot V-Wing",
	},
	"TWI210": {
		ID: "TWI210",
		Name: "Lux Bonteri",
	},
	"TWI211": {
		ID: "TWI211",
		Name: "Sly Moore",
	},
	"TWI212": {
		ID: "TWI212",
		Name: "Freelance Assassin",
	},
	"TWI213": {
		ID: "TWI213",
		Name: "Sanctioner's Shuttle",
	},
	"TWI214": {
		ID: "TWI214",
		Name: "Hidden Sharpshooter",
	},
	"TWI215": {
		ID: "TWI215",
		Name: "Geonosis Patrol Fighter",
	},
	"TWI216": {
		ID: "TWI216",
		Name: "Fives",
	},
	"TWI217": {
		ID: "TWI217",
		Name: "Tri-Droid Suppressor",
	},
	"TWI218": {
		ID: "TWI218",
		Name: "Droid Cohort",
	},
	"TWI219": {
		ID: "TWI219",
		Name: "On Top of Things",
	},
	"TWI220": {
		ID: "TWI220",
		Name: "Shadowed Intentions",
	},
	"TWI221": {
		ID: "TWI221",
		Name: "In Pursuit",
	},
	"TWI222": {
		ID: "TWI222",
		Name: "Political Pressure",
	},
	"TWI223": {
		ID: "TWI223",
		Name: "Unmasking the Conspiracy",
	},
	"TWI224": {
		ID: "TWI224",
		Name: "Breaking In",
	},
	"TWI225": {
		ID: "TWI225",
		Name: "Now There Are Two of Them",
	},
	"TWI226": {
		ID: "TWI226",
		Name: "Waylay",
	},
	"TWI227": {
		ID: "TWI227",
		Name: "Prisoner of War",
	},
	"TWI228": {
		ID: "TWI228",
		Name: "Droid Starfighter",
	},
	"TWI229": {
		ID: "TWI229",
		Name: "Battle Droid Escort",
	},
	"TWI230": {
		ID: "TWI230",
		Name: "Super Battle Droid",
	},
	"TWI231": {
		ID: "TWI231",
		Name: "Dwarf Spider Droid",
	},
	"TWI232": {
		ID: "TWI232",
		Name: "Patrolling AAT",
	},
	"TWI233": {
		ID: "TWI233",
		Name: "Hailfire Tank",
	},
	"TWI234": {
		ID: "TWI234",
		Name: "The Invisible Hand",
	},
	"TWI235": {
		ID: "TWI235",
		Name: "Battle Droid Legion",
	},
	"TWI236": {
		ID: "TWI236",
		Name: "Grievous's Wheel BIke",
	},
	"TWI237": {
		ID: "TWI237",
		Name: "Droid Deployment",
	},
	"TWI238": {
		ID: "TWI238",
		Name: "Merciless Contest",
	},
	"TWI239": {
		ID: "TWI239",
		Name: "Execute Order 66",
	},
	"TWI240": {
		ID: "TWI240",
		Name: "332nd Stalwart",
	},
	"TWI241": {
		ID: "TWI241",
		Name: "Phase I Clone Trooper",
	},
	"TWI242": {
		ID: "TWI242",
		Name: "Phase II Clone Trooper",
	},
	"TWI243": {
		ID: "TWI243",
		Name: "Republic Commando",
	},
	"TWI244": {
		ID: "TWI244",
		Name: "Eta-2 Light Interceptor",
	},
	"TWI245": {
		ID: "TWI245",
		Name: "Armored Saber Tank",
	},
	"TWI246": {
		ID: "TWI246",
		Name: "Tranquility",
	},
	"TWI247": {
		ID: "TWI247",
		Name: "AT-TE Vanguard",
	},
	"TWI248": {
		ID: "TWI248",
		Name: "Ahsoka's Padawan Lightsaber",
	},
	"TWI249": {
		ID: "TWI249",
		Name: "Heroes on Both Sides",
	},
	"TWI250": {
		ID: "TWI250",
		Name: "Sword and Shield Maneuver",
	},
	"TWI251": {
		ID: "TWI251",
		Name: "Drop In",
	},
	"TWI252": {
		ID: "TWI252",
		Name: "Aggrieved Parliamentarian",
	},
	"TWI253": {
		ID: "TWI253",
		Name: "Headhunter Squadron",
	},
	"TWI254": {
		ID: "TWI254",
		Name: "Volunteer Soldier",
	},
	"TWI255": {
		ID: "TWI255",
		Name: "Brain Invaders",
	},
	"TWI256": {
		ID: "TWI256",
		Name: "Hold-Out Blaster",
	},
	"TWI257": {
		ID: "TWI257",
		Name: "Private Manufacturing",
	},
	"JTL001": {
		ID: "JTL001",
		Name: "Asajj Ventress",
	},
	"JTL002": {
		ID: "JTL002",
		Name: "Grand Admiral Thrawn",
	},
	"JTL003": {
		ID: "JTL003",
		Name: "Lando Calrissian",
	},
	"JTL004": {
		ID: "JTL004",
		Name: "Rose Tico",
	},
	"JTL005": {
		ID: "JTL005",
		Name: "Admiral Piett",
	},
	"JTL006": {
		ID: "JTL006",
		Name: "Darth Vader",
	},
	"JTL007": {
		ID: "JTL007",
		Name: "Admiral Holdo",
	},
	"JTL008": {
		ID: "JTL008",
		Name: "Wedge Antilles",
	},
	"JTL009": {
		ID: "JTL009",
		Name: "Boba Fett",
	},
	"JTL010": {
		ID: "JTL010",
		Name: "Captain Phasma",
	},
	"JTL011": {
		ID: "JTL011",
		Name: "Major Vonreg",
	},
	"JTL012": {
		ID: "JTL012",
		Name: "Luke Skywalker",
	},
	"JTL013": {
		ID: "JTL013",
		Name: "Poe Dameron",
	},
	"JTL014": {
		ID: "JTL014",
		Name: "Admiral Trench",
	},
	"JTL015": {
		ID: "JTL015",
		Name: "Rio Durant",
	},
	"JTL016": {
		ID: "JTL016",
		Name: "Admiral Ackbar",
	},
	"JTL017": {
		ID: "JTL017",
		Name: "Han Solo",
	},
	"JTL018": {
		ID: "JTL018",
		Name: "Kazuda Xiono",
	},
	"JTL019": {
		ID: "JTL019",
		Name: "City in the Clouds",
	},
	"JTL020": {
		ID: "JTL020",
		Name: "Shield Generator Complex",
	},
	"JTL021": {
		ID: "JTL021",
		Name: "Colossus",
	},
	"JTL022": {
		ID: "JTL022",
		Name: "Resistance Headquarters",
	},
	"JTL023": {
		ID: "JTL023",
		Name: "Theed Palace",
	},
	"JTL024": {
		ID: "JTL024",
		Name: "Data Vault",
	},
	"JTL025": {
		ID: "JTL025",
		Name: "Thermal Oscillator",
	},
	"JTL026": {
		ID: "JTL026",
		Name: "Massassi Temple",
	},
	"JTL027": {
		ID: "JTL027",
		Name: "Nadiri Dockyards",
	},
	"JTL028": {
		ID: "JTL028",
		Name: "Nabat Village",
	},
	"JTL029": {
		ID: "JTL029",
		Name: "Chopper Base",
	},
	"JTL030": {
		ID: "JTL030",
		Name: "Mos Eisley",
	},
	"JTL031": {
		ID: "JTL031",
		Name: "Lake Country",
	},
	"JTL032": {
		ID: "JTL032",
		Name: "Director Krennic",
	},
	"JTL033": {
		ID: "JTL033",
		Name: "Onyx Squadron Brute",
	},
	"JTL034": {
		ID: "JTL034",
		Name: "Interceptor Ace",
	},
	"JTL035": {
		ID: "JTL035",
		Name: "Tam Ryvora",
	},
	"JTL036": {
		ID: "JTL036",
		Name: "Iden Versio",
	},
	"JTL037": {
		ID: "JTL037",
		Name: "Banshee",
	},
	"JTL038": {
		ID: "JTL038",
		Name: "Corvus",
	},
	"JTL039": {
		ID: "JTL039",
		Name: "Chimaera",
	},
	"JTL040": {
		ID: "JTL040",
		Name: "Fleet Interdictor",
	},
	"JTL041": {
		ID: "JTL041",
		Name: "Annihilator",
	},
	"JTL042": {
		ID: "JTL042",
		Name: "Power from Pain",
	},
	"JTL043": {
		ID: "JTL043",
		Name: "No Glory, Only Results",
	},
	"JTL044": {
		ID: "JTL044",
		Name: "Echo Base Engineer",
	},
	"JTL045": {
		ID: "JTL045",
		Name: "Hera Syndulla",
	},
	"JTL046": {
		ID: "JTL046",
		Name: "Paige Tico",
	},
	"JTL047": {
		ID: "JTL047",
		Name: "Admiral Yularen",
	},
	"JTL048": {
		ID: "JTL048",
		Name: "Cassian Andor",
	},
	"JTL049": {
		ID: "JTL049",
		Name: "L3-37",
	},
	"JTL050": {
		ID: "JTL050",
		Name: "Phantom II",
	},
	"JTL051": {
		ID: "JTL051",
		Name: "Red Squadron X-Wing",
	},
	"JTL052": {
		ID: "JTL052",
		Name: "D'Qar Cargo Frigate",
	},
	"JTL053": {
		ID: "JTL053",
		Name: "The Ghost",
	},
	"JTL054": {
		ID: "JTL054",
		Name: "Gold Leader",
	},
	"JTL055": {
		ID: "JTL055",
		Name: "You're All Clear, Kid",
	},
	"JTL056": {
		ID: "JTL056",
		Name: "Hondo Ohnaka",
	},
	"JTL057": {
		ID: "JTL057",
		Name: "Astromech Pilot",
	},
	"JTL058": {
		ID: "JTL058",
		Name: "Academy Graduate",
	},
	"JTL059": {
		ID: "JTL059",
		Name: "Corporate Defense Shuttle",
	},
	"JTL060": {
		ID: "JTL060",
		Name: "Desperate Commando",
	},
	"JTL061": {
		ID: "JTL061",
		Name: "Royal Security Fighter",
	},
	"JTL062": {
		ID: "JTL062",
		Name: "Silver Angel",
	},
	"JTL063": {
		ID: "JTL063",
		Name: "Landing Shuttle",
	},
	"JTL064": {
		ID: "JTL064",
		Name: "Omicron Strike Craft",
	},
	"JTL065": {
		ID: "JTL065",
		Name: "Outer Rim Outlaws",
	},
	"JTL066": {
		ID: "JTL066",
		Name: "Trace Martez",
	},
	"JTL067": {
		ID: "JTL067",
		Name: "Cloaked StarViper",
	},
	"JTL068": {
		ID: "JTL068",
		Name: "Perimeter AT-RT",
	},
	"JTL069": {
		ID: "JTL069",
		Name: "Munificent Frigate",
	},
	"JTL070": {
		ID: "JTL070",
		Name: "U-Wing Lander",
	},
	"JTL071": {
		ID: "JTL071",
		Name: "CR90 Relief Runner",
	},
	"JTL072": {
		ID: "JTL072",
		Name: "Wing Guard Security Team",
	},
	"JTL073": {
		ID: "JTL073",
		Name: "Grim Valor",
	},
	"JTL074": {
		ID: "JTL074",
		Name: "Close the Shield Gate",
	},
	"JTL075": {
		ID: "JTL075",
		Name: "Repair",
	},
	"JTL076": {
		ID: "JTL076",
		Name: "Covering the Wing",
	},
	"JTL077": {
		ID: "JTL077",
		Name: "In the Heat of Battle",
	},
	"JTL078": {
		ID: "JTL078",
		Name: "Direct Hit",
	},
	"JTL079": {
		ID: "JTL079",
		Name: "Out the Airlock",
	},
	"JTL080": {
		ID: "JTL080",
		Name: "Nebula Ignition",
	},
	"JTL081": {
		ID: "JTL081",
		Name: "First Order TIE Fighter",
	},
	"JTL082": {
		ID: "JTL082",
		Name: "Kijimi Patrollers",
	},
	"JTL083": {
		ID: "JTL083",
		Name: "Pantoran Starship Thief",
	},
	"JTL084": {
		ID: "JTL084",
		Name: "Wingman Victor Two",
	},
	"JTL085": {
		ID: "JTL085",
		Name: "Victor Leader",
	},
	"JTL086": {
		ID: "JTL086",
		Name: "Wingman Victor Three",
	},
	"JTL087": {
		ID: "JTL087",
		Name: "TIE Ambush Squadron",
	},
	"JTL088": {
		ID: "JTL088",
		Name: "Captain Phasma",
	},
	"JTL089": {
		ID: "JTL089",
		Name: "The Invisible Hand",
	},
	"JTL090": {
		ID: "JTL090",
		Name: "Executor",
	},
	"JTL091": {
		ID: "JTL091",
		Name: "Apology Accepted",
	},
	"JTL092": {
		ID: "JTL092",
		Name: "Scramble Fighters",
	},
	"JTL093": {
		ID: "JTL093",
		Name: "Nien Nunb",
	},
	"JTL094": {
		ID: "JTL094",
		Name: "Luke Skywalker",
	},
	"JTL095": {
		ID: "JTL095",
		Name: "Phoenix Squadron A-Wing",
	},
	"JTL096": {
		ID: "JTL096",
		Name: "Blue Leader",
	},
	"JTL097": {
		ID: "JTL097",
		Name: "Leia Organa",
	},
	"JTL098": {
		ID: "JTL098",
		Name: "Snap Wexley",
	},
	"JTL099": {
		ID: "JTL099",
		Name: "Veteran Fleet Officer",
	},
	"JTL100": {
		ID: "JTL100",
		Name: "Poe Dameron",
	},
	"JTL101": {
		ID: "JTL101",
		Name: "Red Leader",
	},
	"JTL102": {
		ID: "JTL102",
		Name: "Resistance Blue Squadron",
	},
	"JTL103": {
		ID: "JTL103",
		Name: "Chewbacca",
	},
	"JTL104": {
		ID: "JTL104",
		Name: "Raddus",
	},
	"JTL105": {
		ID: "JTL105",
		Name: "The Starhawk",
	},
	"JTL106": {
		ID: "JTL106",
		Name: "Unity of Purpose",
	},
	"JTL107": {
		ID: "JTL107",
		Name: "Bunker Defender",
	},
	"JTL108": {
		ID: "JTL108",
		Name: "Clone Pilot",
	},
	"JTL109": {
		ID: "JTL109",
		Name: "Jarek Yeager",
	},
	"JTL110": {
		ID: "JTL110",
		Name: "Scouting Headhunter",
	},
	"JTL111": {
		ID: "JTL111",
		Name: "Seasoned Fleet Admiral",
	},
	"JTL112": {
		ID: "JTL112",
		Name: "Eager Escort Fighter",
	},
	"JTL113": {
		ID: "JTL113",
		Name: "Homestead Militia",
	},
	"JTL114": {
		ID: "JTL114",
		Name: "Adept ARC-170",
	},
	"JTL115": {
		ID: "JTL115",
		Name: "Clone Combat Squadron",
	},
	"JTL116": {
		ID: "JTL116",
		Name: "Dornean Gunship",
	},
	"JTL117": {
		ID: "JTL117",
		Name: "General Draven",
	},
	"JTL118": {
		ID: "JTL118",
		Name: "MC30 Assault Frigate",
	},
	"JTL119": {
		ID: "JTL119",
		Name: "Resupply Carrier",
	},
	"JTL120": {
		ID: "JTL120",
		Name: "Dorsal Turret",
	},
	"JTL121": {
		ID: "JTL121",
		Name: "Salvage",
	},
	"JTL122": {
		ID: "JTL122",
		Name: "All Wings Report In",
	},
	"JTL123": {
		ID: "JTL123",
		Name: "Dogfight",
	},
	"JTL124": {
		ID: "JTL124",
		Name: "Tandem Assault",
	},
	"JTL125": {
		ID: "JTL125",
		Name: "Air Superiority",
	},
	"JTL126": {
		ID: "JTL126",
		Name: "Eject",
	},
	"JTL127": {
		ID: "JTL127",
		Name: "Lightspeed Assault",
	},
	"JTL128": {
		ID: "JTL128",
		Name: "Prepare for Takeoff",
	},
	"JTL129": {
		ID: "JTL129",
		Name: "Focus Fire",
	},
	"JTL130": {
		ID: "JTL130",
		Name: "Timely Reinforcements",
	},
	"JTL131": {
		ID: "JTL131",
		Name: "Turbolaser Salvo",
	},
	"JTL132": {
		ID: "JTL132",
		Name: "First Order Stormtrooper",
	},
	"JTL133": {
		ID: "JTL133",
		Name: "Allegiant General Pryde",
	},
	"JTL134": {
		ID: "JTL134",
		Name: "General Hux",
	},
	"JTL135": {
		ID: "JTL135",
		Name: "Special Forces TIE Fighter",
	},
	"JTL136": {
		ID: "JTL136",
		Name: "Prototype TIE Advanced",
	},
	"JTL137": {
		ID: "JTL137",
		Name: "Vonreg's TIE Interceptor",
	},
	"JTL138": {
		ID: "JTL138",
		Name: "Decimator of Dissidents",
	},
	"JTL139": {
		ID: "JTL139",
		Name: "Dengar",
	},
	"JTL140": {
		ID: "JTL140",
		Name: "IG-2000",
	},
	"JTL141": {
		ID: "JTL141",
		Name: "IG-88",
	},
	"JTL142": {
		ID: "JTL142",
		Name: "Darth Vader",
	},
	"JTL143": {
		ID: "JTL143",
		Name: "Devastator",
	},
	"JTL144": {
		ID: "JTL144",
		Name: "No Disintegrations",
	},
	"JTL145": {
		ID: "JTL145",
		Name: "BB-8",
	},
	"JTL146": {
		ID: "JTL146",
		Name: "Massassi Tactical Officer",
	},
	"JTL147": {
		ID: "JTL147",
		Name: "Black One",
	},
	"JTL148": {
		ID: "JTL148",
		Name: "Frisk",
	},
	"JTL149": {
		ID: "JTL149",
		Name: "Red Squadron Y-Wing",
	},
	"JTL150": {
		ID: "JTL150",
		Name: "Biggs Darklighter",
	},
	"JTL151": {
		ID: "JTL151",
		Name: "Red Five",
	},
	"JTL152": {
		ID: "JTL152",
		Name: "Tactical Heavy Bomber",
	},
	"JTL153": {
		ID: "JTL153",
		Name: "Rebellious Hammerhead",
	},
	"JTL154": {
		ID: "JTL154",
		Name: "Profundity",
	},
	"JTL155": {
		ID: "JTL155",
		Name: "They Hate That Ship",
	},
	"JTL156": {
		ID: "JTL156",
		Name: "Trench Run",
	},
	"JTL157": {
		ID: "JTL157",
		Name: "Relentless Firespray",
	},
	"JTL158": {
		ID: "JTL158",
		Name: "Crackshot V-Wing",
	},
	"JTL159": {
		ID: "JTL159",
		Name: "Determined Recruit",
	},
	"JTL160": {
		ID: "JTL160",
		Name: "Supporting Eta-2",
	},
	"JTL161": {
		ID: "JTL161",
		Name: "Captain Tarkin",
	},
	"JTL162": {
		ID: "JTL162",
		Name: "Droid Missile Platform",
	},
	"JTL163": {
		ID: "JTL163",
		Name: "AT-DP Occupier",
	},
	"JTL164": {
		ID: "JTL164",
		Name: "Cham Syndulla",
	},
	"JTL165": {
		ID: "JTL165",
		Name: "Hunting Aggressor",
	},
	"JTL166": {
		ID: "JTL166",
		Name: "Orbiting K-Wing",
	},
	"JTL167": {
		ID: "JTL167",
		Name: "Occupier Siege Tank",
	},
	"JTL168": {
		ID: "JTL168",
		Name: "Insurgent Saboteurs",
	},
	"JTL169": {
		ID: "JTL169",
		Name: "Shadow Caster",
	},
	"JTL170": {
		ID: "JTL170",
		Name: "War Juggernaut",
	},
	"JTL171": {
		ID: "JTL171",
		Name: "Targeting Computer",
	},
	"JTL172": {
		ID: "JTL172",
		Name: "Twin Laser Turret",
	},
	"JTL173": {
		ID: "JTL173",
		Name: "Fight Fire With Fire",
	},
	"JTL174": {
		ID: "JTL174",
		Name: "Hotshot Maneuver",
	},
	"JTL175": {
		ID: "JTL175",
		Name: "System Shock",
	},
	"JTL176": {
		ID: "JTL176",
		Name: "Shoot Down",
	},
	"JTL177": {
		ID: "JTL177",
		Name: "Stay on Target",
	},
	"JTL178": {
		ID: "JTL178",
		Name: "Face Off",
	},
	"JTL179": {
		ID: "JTL179",
		Name: "Koiogran Turn",
	},
	"JTL180": {
		ID: "JTL180",
		Name: "Piercing Shot",
	},
	"JTL181": {
		ID: "JTL181",
		Name: "Planetary Bombardment",
	},
	"JTL182": {
		ID: "JTL182",
		Name: "Rampart",
	},
	"JTL183": {
		ID: "JTL183",
		Name: "Zygerrian Starhopper",
	},
	"JTL184": {
		ID: "JTL184",
		Name: "Contracted Jumpmaster",
	},
	"JTL185": {
		ID: "JTL185",
		Name: "Hound's Tooth",
	},
	"JTL186": {
		ID: "JTL186",
		Name: "Mist Hunter",
	},
	"JTL187": {
		ID: "JTL187",
		Name: "Bossk",
	},
	"JTL188": {
		ID: "JTL188",
		Name: "Moff Gideon",
	},
	"JTL189": {
		ID: "JTL189",
		Name: "Boba Fett",
	},
	"JTL190": {
		ID: "JTL190",
		Name: "Techno Union Transport",
	},
	"JTL191": {
		ID: "JTL191",
		Name: "Invincible",
	},
	"JTL192": {
		ID: "JTL192",
		Name: "In Debt to Crimson Dawn",
	},
	"JTL193": {
		ID: "JTL193",
		Name: "I Have You Now",
	},
	"JTL194": {
		ID: "JTL194",
		Name: "Heartless Tactics",
	},
	"JTL195": {
		ID: "JTL195",
		Name: "Cat and Mouse",
	},
	"JTL196": {
		ID: "JTL196",
		Name: "Dagger Squadron Pilot",
	},
	"JTL197": {
		ID: "JTL197",
		Name: "Anakin Skywalker",
	},
	"JTL198": {
		ID: "JTL198",
		Name: "Fireball",
	},
	"JTL199": {
		ID: "JTL199",
		Name: "Blade Squadron B-Wing",
	},
	"JTL200": {
		ID: "JTL200",
		Name: "Shuttle Tydirium",
	},
	"JTL201": {
		ID: "JTL201",
		Name: "Ahsoka Tano",
	},
	"JTL202": {
		ID: "JTL202",
		Name: "Black Squadron Scout Wing",
	},
	"JTL203": {
		ID: "JTL203",
		Name: "Han Solo",
	},
	"JTL204": {
		ID: "JTL204",
		Name: "Home One",
	},
	"JTL205": {
		ID: "JTL205",
		Name: "Commence Patrol",
	},
	"JTL206": {
		ID: "JTL206",
		Name: "Fly Casual",
	},
	"JTL207": {
		ID: "JTL207",
		Name: "Jam Communications",
	},
	"JTL208": {
		ID: "JTL208",
		Name: "Never Tell Me The Odds",
	},
	"JTL209": {
		ID: "JTL209",
		Name: "It's a Trap",
	},
	"JTL210": {
		ID: "JTL210",
		Name: "The Mandalorian",
	},
	"JTL211": {
		ID: "JTL211",
		Name: "Independent Smuggler",
	},
	"JTL212": {
		ID: "JTL212",
		Name: "Republic Y-Wing",
	},
	"JTL213": {
		ID: "JTL213",
		Name: "Sidon Ithano",
	},
	"JTL214": {
		ID: "JTL214",
		Name: "X-34 Landspeeder",
	},
	"JTL215": {
		ID: "JTL215",
		Name: "BoShek",
	},
	"JTL216": {
		ID: "JTL216",
		Name: "Contracted Hunter",
	},
	"JTL217": {
		ID: "JTL217",
		Name: "Death Space Skirmisher",
	},
	"JTL218": {
		ID: "JTL218",
		Name: "Guerilla Soldier",
	},
	"JTL219": {
		ID: "JTL219",
		Name: "Rafa Martez",
	},
	"JTL220": {
		ID: "JTL220",
		Name: "Skyway Cloud Car",
	},
	"JTL221": {
		ID: "JTL221",
		Name: "Stolen AT-Hauler",
	},
	"JTL222": {
		ID: "JTL222",
		Name: "Kimogila Heavy Fighter",
	},
	"JTL223": {
		ID: "JTL223",
		Name: "Razor Crest",
	},
	"JTL224": {
		ID: "JTL224",
		Name: "Shadowed Hover Tank",
	},
	"JTL225": {
		ID: "JTL225",
		Name: "Corporate Light Cruiser",
	},
	"JTL226": {
		ID: "JTL226",
		Name: "Radiant VII",
	},
	"JTL227": {
		ID: "JTL227",
		Name: "Superheavy Ion Cannon",
	},
	"JTL228": {
		ID: "JTL228",
		Name: "Barrel Roll",
	},
	"JTL229": {
		ID: "JTL229",
		Name: "Diversion",
	},
	"JTL230": {
		ID: "JTL230",
		Name: "Electromagnetic Pulse",
	},
	"JTL231": {
		ID: "JTL231",
		Name: "Punch It",
	},
	"JTL232": {
		ID: "JTL232",
		Name: "Jump to Lightspeed",
	},
	"JTL233": {
		ID: "JTL233",
		Name: "Sweep the Area",
	},
	"JTL234": {
		ID: "JTL234",
		Name: "Torpedo Barrage",
	},
	"JTL235": {
		ID: "JTL235",
		Name: "Commandeer",
	},
	"JTL236": {
		ID: "JTL236",
		Name: "Indoctrinated Conscript",
	},
	"JTL237": {
		ID: "JTL237",
		Name: "TIE Bomber",
	},
	"JTL238": {
		ID: "JTL238",
		Name: "Sith Trooper",
	},
	"JTL239": {
		ID: "JTL239",
		Name: "TIE Dagger Vanguard",
	},
	"JTL240": {
		ID: "JTL240",
		Name: "Fett's Firespray",
	},
	"JTL241": {
		ID: "JTL241",
		Name: "Rogue-class Starfighter",
	},
	"JTL242": {
		ID: "JTL242",
		Name: "Shuttle ST-149",
	},
	"JTL243": {
		ID: "JTL243",
		Name: "Quasar TIE Carrier",
	},
	"JTL244": {
		ID: "JTL244",
		Name: "There Is No Escape",
	},
	"JTL245": {
		ID: "JTL245",
		Name: "R2-D2",
	},
	"JTL246": {
		ID: "JTL246",
		Name: "Hopeful Volunteer",
	},
	"JTL247": {
		ID: "JTL247",
		Name: "Resistance X-Wing",
	},
	"JTL248": {
		ID: "JTL248",
		Name: "Dilapidated Ski Speeder",
	},
	"JTL249": {
		ID: "JTL249",
		Name: "Millennium Falcon",
	},
	"JTL250": {
		ID: "JTL250",
		Name: "Sabine's Masterpiece",
	},
	"JTL251": {
		ID: "JTL251",
		Name: "Jedi Light Cruiser",
	},
	"JTL252": {
		ID: "JTL252",
		Name: "Tantive IV",
	},
	"JTL253": {
		ID: "JTL253",
		Name: "Coordinated Front",
	},
	"JTL254": {
		ID: "JTL254",
		Name: "Dedicated Wingmen",
	},
	"JTL255": {
		ID: "JTL255",
		Name: "Sullustan Spacer",
	},
	"JTL256": {
		ID: "JTL256",
		Name: "Swarming Vulture Droid",
	},
	"JTL257": {
		ID: "JTL257",
		Name: "Flanking Fang Fighter",
	},
	"JTL258": {
		ID: "JTL258",
		Name: "Corellian Freighter",
	},
	"JTL259": {
		ID: "JTL259",
		Name: "Retrofitted Airspeeder",
	},
	"JTL260": {
		ID: "JTL260",
		Name: "Death Star Plans",
	},
	"JTL261": {
		ID: "JTL261",
		Name: "Attack Run",
	},
	"JTL262": {
		ID: "JTL262",
		Name: "Evasive Maneuver",
	},
	"LOF001": {
		ID: "LOF001",
		Name: "Kylo Ren",
	},
	"LOF002": {
		ID: "LOF002",
		Name: "Mother Talzin",
	},
	"LOF003": {
		ID: "LOF003",
		Name: "Ahsoka Tano",
	},
	"LOF004": {
		ID: "LOF004",
		Name: "Kanan Jarrus",
	},
	"LOF005": {
		ID: "LOF005",
		Name: "Morgan Elsbeth",
	},
	"LOF006": {
		ID: "LOF006",
		Name: "Supreme Leader Snoke",
	},
	"LOF007": {
		ID: "LOF007",
		Name: "Avar Kriss",
	},
	"LOF008": {
		ID: "LOF008",
		Name: "Obi-Wan Kenobi",
	},
	"LOF009": {
		ID: "LOF009",
		Name: "Darth Maul",
	},
	"LOF010": {
		ID: "LOF010",
		Name: "Third Sister",
	},
	"LOF011": {
		ID: "LOF011",
		Name: "Kit Fisto",
	},
	"LOF012": {
		ID: "LOF012",
		Name: "Rey",
	},
	"LOF013": {
		ID: "LOF013",
		Name: "Barriss Offee",
	},
	"LOF014": {
		ID: "LOF014",
		Name: "Grand Inquisitor",
	},
	"LOF015": {
		ID: "LOF015",
		Name: "Cal Kestis",
	},
	"LOF016": {
		ID: "LOF016",
		Name: "Qui-Gon Jinn",
	},
	"LOF017": {
		ID: "LOF017",
		Name: "Darth Revan",
	},
	"LOF018": {
		ID: "LOF018",
		Name: "Anakin Skywalker",
	},
	"LOF019": {
		ID: "LOF019",
		Name: "Vergence Temple",
	},
	"LOF020": {
		ID: "LOF020",
		Name: "Nightsister Lair",
	},
	"LOF021": {
		ID: "LOF021",
		Name: "Shadowed Undercity",
	},
	"LOF022": {
		ID: "LOF022",
		Name: "Mystic Monastery",
	},
	"LOF023": {
		ID: "LOF023",
		Name: "Jedi Temple",
	},
	"LOF024": {
		ID: "LOF024",
		Name: "Starlight Temple",
	},
	"LOF025": {
		ID: "LOF025",
		Name: "Temple of Destruction",
	},
	"LOF026": {
		ID: "LOF026",
		Name: "Fortress Vader",
	},
	"LOF027": {
		ID: "LOF027",
		Name: "Strangled Cliffs",
	},
	"LOF028": {
		ID: "LOF028",
		Name: "Tomb of Eilram",
	},
	"LOF029": {
		ID: "LOF029",
		Name: "Crystal Caves",
	},
	"LOF030": {
		ID: "LOF030",
		Name: "The Holy City",
	},
	"LOF031": {
		ID: "LOF031",
		Name: "Karis",
	},
	"LOF032": {
		ID: "LOF032",
		Name: "Magistrate's Scout",
	},
	"LOF033": {
		ID: "LOF033",
		Name: "Nameless Terror",
	},
	"LOF034": {
		ID: "LOF034",
		Name: "Supremacy TIE/sf",
	},
	"LOF035": {
		ID: "LOF035",
		Name: "Talzin's Assassin",
	},
	"LOF036": {
		ID: "LOF036",
		Name: "Old Daka",
	},
	"LOF037": {
		ID: "LOF037",
		Name: "Darth Vader",
	},
	"LOF038": {
		ID: "LOF038",
		Name: "Pong Krell",
	},
	"LOF039": {
		ID: "LOF039",
		Name: "Darth Sidious",
	},
	"LOF040": {
		ID: "LOF040",
		Name: "Kylo Ren's Lightsaber",
	},
	"LOF041": {
		ID: "LOF041",
		Name: "Drain Essence",
	},
	"LOF042": {
		ID: "LOF042",
		Name: "Always Two",
	},
	"LOF043": {
		ID: "LOF043",
		Name: "The Tragedy of Plagueis",
	},
	"LOF044": {
		ID: "LOF044",
		Name: "Loth-Wolf",
	},
	"LOF045": {
		ID: "LOF045",
		Name: "Yaddle",
	},
	"LOF046": {
		ID: "LOF046",
		Name: "Ezra Bridger",
	},
	"LOF047": {
		ID: "LOF047",
		Name: "T-6 Shuttle 1974",
	},
	"LOF048": {
		ID: "LOF048",
		Name: "Itinerant Warrior",
	},
	"LOF049": {
		ID: "LOF049",
		Name: "Jedi Guardian",
	},
	"LOF050": {
		ID: "LOF050",
		Name: "Plo Koon",
	},
	"LOF051": {
		ID: "LOF051",
		Name: "Jedi Holocron",
	},
	"LOF052": {
		ID: "LOF052",
		Name: "Jedi Trials",
	},
	"LOF053": {
		ID: "LOF053",
		Name: "Heirloom Lightsaber",
	},
	"LOF054": {
		ID: "LOF054",
		Name: "Calm in the Storm",
	},
	"LOF055": {
		ID: "LOF055",
		Name: "Dume",
	},
	"LOF056": {
		ID: "LOF056",
		Name: "Size Matters Not",
	},
	"LOF057": {
		ID: "LOF057",
		Name: "Owen Lars",
	},
	"LOF058": {
		ID: "LOF058",
		Name: "Guardian of the Whills",
	},
	"LOF059": {
		ID: "LOF059",
		Name: "Nightsister Warrior",
	},
	"LOF060": {
		ID: "LOF060",
		Name: "Padawan Starfighter",
	},
	"LOF061": {
		ID: "LOF061",
		Name: "Secretive Sage",
	},
	"LOF062": {
		ID: "LOF062",
		Name: "Axe Woves",
	},
	"LOF063": {
		ID: "LOF063",
		Name: "Oggdo Bogdo",
	},
	"LOF064": {
		ID: "LOF064",
		Name: "Tauntaun",
	},
	"LOF065": {
		ID: "LOF065",
		Name: "Watto",
	},
	"LOF066": {
		ID: "LOF066",
		Name: "Awakened Specters",
	},
	"LOF067": {
		ID: "LOF067",
		Name: "Chirrut Îmwe",
	},
	"LOF068": {
		ID: "LOF068",
		Name: "Luthen Rael",
	},
	"LOF069": {
		ID: "LOF069",
		Name: "Graceful Purrgil",
	},
	"LOF070": {
		ID: "LOF070",
		Name: "Anakin Skywalker",
	},
	"LOF071": {
		ID: "LOF071",
		Name: "Grappling Guardian",
	},
	"LOF072": {
		ID: "LOF072",
		Name: "Priestesses of the Force",
	},
	"LOF073": {
		ID: "LOF073",
		Name: "Mythosaur",
	},
	"LOF074": {
		ID: "LOF074",
		Name: "Bolstered Endurance",
	},
	"LOF075": {
		ID: "LOF075",
		Name: "Cure Wounds",
	},
	"LOF076": {
		ID: "LOF076",
		Name: "Soresu Stance",
	},
	"LOF077": {
		ID: "LOF077",
		Name: "Crushing Blow",
	},
	"LOF078": {
		ID: "LOF078",
		Name: "Whirlwind of Power",
	},
	"LOF079": {
		ID: "LOF079",
		Name: "Shatterpoint",
	},
	"LOF080": {
		ID: "LOF080",
		Name: "Exegol Patroller",
	},
	"LOF081": {
		ID: "LOF081",
		Name: "Sith Legionnaire",
	},
	"LOF082": {
		ID: "LOF082",
		Name: "Vaneé",
	},
	"LOF083": {
		ID: "LOF083",
		Name: "Captain Enoch",
	},
	"LOF084": {
		ID: "LOF084",
		Name: "Knight of Ren",
	},
	"LOF085": {
		ID: "LOF085",
		Name: "Praetorian Guard",
	},
	"LOF086": {
		ID: "LOF086",
		Name: "Drengir Spawn",
	},
	"LOF087": {
		ID: "LOF087",
		Name: "Eighth Brother",
	},
	"LOF088": {
		ID: "LOF088",
		Name: "Eye of Sion",
	},
	"LOF089": {
		ID: "LOF089",
		Name: "Supremacy",
	},
	"LOF090": {
		ID: "LOF090",
		Name: "Inquisitor's Lightsaber",
	},
	"LOF091": {
		ID: "LOF091",
		Name: "Craving Power",
	},
	"LOF092": {
		ID: "LOF092",
		Name: "Point Rain Reclaimer",
	},
	"LOF093": {
		ID: "LOF093",
		Name: "Gungi",
	},
	"LOF094": {
		ID: "LOF094",
		Name: "Jedi Consular",
	},
	"LOF095": {
		ID: "LOF095",
		Name: "Lor San Tekka",
	},
	"LOF096": {
		ID: "LOF096",
		Name: "Obi-Wan Kenobi",
	},
	"LOF097": {
		ID: "LOF097",
		Name: "Eeth Koth",
	},
	"LOF098": {
		ID: "LOF098",
		Name: "Leia Organa",
	},
	"LOF099": {
		ID: "LOF099",
		Name: "Paladin Training Corvette",
	},
	"LOF100": {
		ID: "LOF100",
		Name: "Kelleran Beq",
	},
	"LOF101": {
		ID: "LOF101",
		Name: "Yoda",
	},
	"LOF102": {
		ID: "LOF102",
		Name: "Yoda's Lightsaber",
	},
	"LOF103": {
		ID: "LOF103",
		Name: "Following the Path",
	},
	"LOF104": {
		ID: "LOF104",
		Name: "Luminous Beings",
	},
	"LOF105": {
		ID: "LOF105",
		Name: "Oppo Rancisis",
	},
	"LOF106": {
		ID: "LOF106",
		Name: "Acclamator Assault Ship",
	},
	"LOF107": {
		ID: "LOF107",
		Name: "Village Tender",
	},
	"LOF108": {
		ID: "LOF108",
		Name: "Malakili",
	},
	"LOF109": {
		ID: "LOF109",
		Name: "Mynock",
	},
	"LOF110": {
		ID: "LOF110",
		Name: "Hive Defense Wing",
	},
	"LOF111": {
		ID: "LOF111",
		Name: "Maz Kanata",
	},
	"LOF112": {
		ID: "LOF112",
		Name: "Outer Rim Mystic",
	},
	"LOF113": {
		ID: "LOF113",
		Name: "Jedi Temple Guards",
	},
	"LOF114": {
		ID: "LOF114",
		Name: "Kaadu",
	},
	"LOF115": {
		ID: "LOF115",
		Name: "Dagoyan Master",
	},
	"LOF116": {
		ID: "LOF116",
		Name: "Relic Scavenger",
	},
	"LOF117": {
		ID: "LOF117",
		Name: "Sifo-Dyas",
	},
	"LOF118": {
		ID: "LOF118",
		Name: "Terentatek",
	},
	"LOF119": {
		ID: "LOF119",
		Name: "Hyperspace Wayfarer",
	},
	"LOF120": {
		ID: "LOF120",
		Name: "Trident Assault Ship",
	},
	"LOF121": {
		ID: "LOF121",
		Name: "The Purrgil King",
	},
	"LOF122": {
		ID: "LOF122",
		Name: "Pillio Star Compass",
	},
	"LOF123": {
		ID: "LOF123",
		Name: "Directed by the Force",
	},
	"LOF124": {
		ID: "LOF124",
		Name: "Niman Strike",
	},
	"LOF125": {
		ID: "LOF125",
		Name: "The Burden of Masters",
	},
	"LOF126": {
		ID: "LOF126",
		Name: "Overpower",
	},
	"LOF127": {
		ID: "LOF127",
		Name: "Rampage",
	},
	"LOF128": {
		ID: "LOF128",
		Name: "Protect the Pod",
	},
	"LOF129": {
		ID: "LOF129",
		Name: "Acolyte of the Beyond",
	},
	"LOF130": {
		ID: "LOF130",
		Name: "HK-47",
	},
	"LOF131": {
		ID: "LOF131",
		Name: "Strikeship",
	},
	"LOF132": {
		ID: "LOF132",
		Name: "Grand Inquisitor",
	},
	"LOF133": {
		ID: "LOF133",
		Name: "Purge Trooper",
	},
	"LOF134": {
		ID: "LOF134",
		Name: "Heavy Missile Gunship",
	},
	"LOF135": {
		ID: "LOF135",
		Name: "Scythe",
	},
	"LOF136": {
		ID: "LOF136",
		Name: "Thralls of the Coven",
	},
	"LOF137": {
		ID: "LOF137",
		Name: "Savage Opress",
	},
	"LOF138": {
		ID: "LOF138",
		Name: "Sith Holocron",
	},
	"LOF139": {
		ID: "LOF139",
		Name: "Battle Fury",
	},
	"LOF140": {
		ID: "LOF140",
		Name: "Darth Maul's Lightsaber",
	},
	"LOF141": {
		ID: "LOF141",
		Name: "Death Field",
	},
	"LOF142": {
		ID: "LOF142",
		Name: "Adi Gallia",
	},
	"LOF143": {
		ID: "LOF143",
		Name: "Attuned Fyrnock",
	},
	"LOF144": {
		ID: "LOF144",
		Name: "Jedi Starfighter",
	},
	"LOF145": {
		ID: "LOF145",
		Name: "Jedi Knight",
	},
	"LOF146": {
		ID: "LOF146",
		Name: "Ki-Adi-Mundi",
	},
	"LOF147": {
		ID: "LOF147",
		Name: "Kit Fisto's Aethersprite",
	},
	"LOF148": {
		ID: "LOF148",
		Name: "Rey",
	},
	"LOF149": {
		ID: "LOF149",
		Name: "Mace Windu",
	},
	"LOF150": {
		ID: "LOF150",
		Name: "Cin Drallig",
	},
	"LOF151": {
		ID: "LOF151",
		Name: "Knight's Saber",
	},
	"LOF152": {
		ID: "LOF152",
		Name: "Focus Determines Reality",
	},
	"LOF153": {
		ID: "LOF153",
		Name: "Paz Vizsla",
	},
	"LOF154": {
		ID: "LOF154",
		Name: "Witch of the Mist",
	},
	"LOF155": {
		ID: "LOF155",
		Name: "DRK-1 Probe Droid",
	},
	"LOF156": {
		ID: "LOF156",
		Name: "Infused Brawler",
	},
	"LOF157": {
		ID: "LOF157",
		Name: "Cartel Interceptor",
	},
	"LOF158": {
		ID: "LOF158",
		Name: "Hyena Bomber",
	},
	"LOF159": {
		ID: "LOF159",
		Name: "Jedi in Hiding",
	},
	"LOF160": {
		ID: "LOF160",
		Name: "Merrin",
	},
	"LOF161": {
		ID: "LOF161",
		Name: "Tuk'ata",
	},
	"LOF162": {
		ID: "LOF162",
		Name: "Hunting Nexu",
	},
	"LOF163": {
		ID: "LOF163",
		Name: "Quinlan Vos",
	},
	"LOF164": {
		ID: "LOF164",
		Name: "Wampa",
	},
	"LOF165": {
		ID: "LOF165",
		Name: "Asajj Ventress",
	},
	"LOF166": {
		ID: "LOF166",
		Name: "Blockade Runner",
	},
	"LOF167": {
		ID: "LOF167",
		Name: "Saesee Tiin",
	},
	"LOF168": {
		ID: "LOF168",
		Name: "Ravenous Rathtar",
	},
	"LOF169": {
		ID: "LOF169",
		Name: "Invasion Control Ship",
	},
	"LOF170": {
		ID: "LOF170",
		Name: "Bendu",
	},
	"LOF171": {
		ID: "LOF171",
		Name: "Heavy Blaster Cannon",
	},
	"LOF172": {
		ID: "LOF172",
		Name: "Sorcerous Blast",
	},
	"LOF173": {
		ID: "LOF173",
		Name: "Unleash Rage",
	},
	"LOF174": {
		ID: "LOF174",
		Name: "Ataru Onslaught",
	},
	"LOF175": {
		ID: "LOF175",
		Name: "Do or Do Not",
	},
	"LOF176": {
		ID: "LOF176",
		Name: "Lightsaber Throw",
	},
	"LOF177": {
		ID: "LOF177",
		Name: "Time of Crisis",
	},
	"LOF178": {
		ID: "LOF178",
		Name: "Adept of Anger",
	},
	"LOF179": {
		ID: "LOF179",
		Name: "Aurra Sing",
	},
	"LOF180": {
		ID: "LOF180",
		Name: "Deceptive Shade",
	},
	"LOF181": {
		ID: "LOF181",
		Name: "Banking Clan Shuttle",
	},
	"LOF182": {
		ID: "LOF182",
		Name: "Nihil Marauder",
	},
	"LOF183": {
		ID: "LOF183",
		Name: "Shin Hati",
	},
	"LOF184": {
		ID: "LOF184",
		Name: "Second Sister",
	},
	"LOF185": {
		ID: "LOF185",
		Name: "Baylan Skoll",
	},
	"LOF186": {
		ID: "LOF186",
		Name: "Marchion Ro",
	},
	"LOF187": {
		ID: "LOF187",
		Name: "Corrupted Saber",
	},
	"LOF188": {
		ID: "LOF188",
		Name: "As I Have Foreseen",
	},
	"LOF189": {
		ID: "LOF189",
		Name: "Liberated by Darkness",
	},
	"LOF190": {
		ID: "LOF190",
		Name: "Anakin Skywalker",
	},
	"LOF191": {
		ID: "LOF191",
		Name: "BD-1",
	},
	"LOF192": {
		ID: "LOF192",
		Name: "N-1 Starfighter",
	},
	"LOF193": {
		ID: "LOF193",
		Name: "Youngling Padawan",
	},
	"LOF194": {
		ID: "LOF194",
		Name: "J-Type Nubian Starship",
	},
	"LOF195": {
		ID: "LOF195",
		Name: "Vernestra Rwoh",
	},
	"LOF196": {
		ID: "LOF196",
		Name: "Jedi Sentinel",
	},
	"LOF197": {
		ID: "LOF197",
		Name: "Qui-Gon Jinn's Aethersprite",
	},
	"LOF198": {
		ID: "LOF198",
		Name: "Stinger Mantis",
	},
	"LOF199": {
		ID: "LOF199",
		Name: "Depa Billaba",
	},
	"LOF200": {
		ID: "LOF200",
		Name: "Qui-Gon Jinn",
	},
	"LOF201": {
		ID: "LOF201",
		Name: "Qui-Gon Jinn's Lightsaber",
	},
	"LOF202": {
		ID: "LOF202",
		Name: "Mind Trick",
	},
	"LOF203": {
		ID: "LOF203",
		Name: "Premonition of Doom",
	},
	"LOF204": {
		ID: "LOF204",
		Name: "Zuckuss",
	},
	"LOF205": {
		ID: "LOF205",
		Name: "Force Speed",
	},
	"LOF206": {
		ID: "LOF206",
		Name: "Babu Frik",
	},
	"LOF207": {
		ID: "LOF207",
		Name: "Loth-Cat",
	},
	"LOF208": {
		ID: "LOF208",
		Name: "Mysterious Hermit",
	},
	"LOF209": {
		ID: "LOF209",
		Name: "Tusken Tracker",
	},
	"LOF210": {
		ID: "LOF210",
		Name: "Charging Phillak",
	},
	"LOF211": {
		ID: "LOF211",
		Name: "Dooku",
	},
	"LOF212": {
		ID: "LOF212",
		Name: "Life Wind Sage",
	},
	"LOF213": {
		ID: "LOF213",
		Name: "The Legacy Run",
	},
	"LOF214": {
		ID: "LOF214",
		Name: "Sorcerers of Tund",
	},
	"LOF215": {
		ID: "LOF215",
		Name: "Ascension Cable",
	},
	"LOF216": {
		ID: "LOF216",
		Name: "Disturbance in the Force",
	},
	"LOF217": {
		ID: "LOF217",
		Name: "Force Slow",
	},
	"LOF218": {
		ID: "LOF218",
		Name: "Impossible Escape",
	},
	"LOF219": {
		ID: "LOF219",
		Name: "Psychometry",
	},
	"LOF220": {
		ID: "LOF220",
		Name: "Shien Flurry",
	},
	"LOF221": {
		ID: "LOF221",
		Name: "Trust Your Instincts",
	},
	"LOF222": {
		ID: "LOF222",
		Name: "A Precarious Predicament",
	},
	"LOF223": {
		ID: "LOF223",
		Name: "Force Illusion",
	},
	"LOF224": {
		ID: "LOF224",
		Name: "Pounce",
	},
	"LOF225": {
		ID: "LOF225",
		Name: "Three Lessons",
	},
	"LOF226": {
		ID: "LOF226",
		Name: "Tip the Scale",
	},
	"LOF227": {
		ID: "LOF227",
		Name: "The Will of the Force",
	},
	"LOF228": {
		ID: "LOF228",
		Name: "Forged Starfighter",
	},
	"LOF229": {
		ID: "LOF229",
		Name: "Kylo Ren",
	},
	"LOF230": {
		ID: "LOF230",
		Name: "Fallen Jedi",
	},
	"LOF231": {
		ID: "LOF231",
		Name: "Darth Tyranus",
	},
	"LOF232": {
		ID: "LOF232",
		Name: "Sandtrooper Cavalry",
	},
	"LOF233": {
		ID: "LOF233",
		Name: "Scimitar",
	},
	"LOF234": {
		ID: "LOF234",
		Name: "Darth Malak",
	},
	"LOF235": {
		ID: "LOF235",
		Name: "HK-87 Assassin Droid",
	},
	"LOF236": {
		ID: "LOF236",
		Name: "Army of the Dead",
	},
	"LOF237": {
		ID: "LOF237",
		Name: "The Son",
	},
	"LOF238": {
		ID: "LOF238",
		Name: "Darth Revan's Lightsabers",
	},
	"LOF239": {
		ID: "LOF239",
		Name: "Consumed by the Dark Side",
	},
	"LOF240": {
		ID: "LOF240",
		Name: "Flight of the Inquisitor",
	},
	"LOF241": {
		ID: "LOF241",
		Name: "In the Shadows",
	},
	"LOF242": {
		ID: "LOF242",
		Name: "Refugee of the Path",
	},
	"LOF243": {
		ID: "LOF243",
		Name: "Caretaker Matron",
	},
	"LOF244": {
		ID: "LOF244",
		Name: "Jedi Vector",
	},
	"LOF245": {
		ID: "LOF245",
		Name: "Vulptex",
	},
	"LOF246": {
		ID: "LOF246",
		Name: "Grogu",
	},
	"LOF247": {
		ID: "LOF247",
		Name: "Gungan Warrior",
	},
	"LOF248": {
		ID: "LOF248",
		Name: "Jocasta Nu",
	},
	"LOF249": {
		ID: "LOF249",
		Name: "Luke Skywalker",
	},
	"LOF250": {
		ID: "LOF250",
		Name: "Medical Frigate",
	},
	"LOF251": {
		ID: "LOF251",
		Name: "Blue Squadron Assault Wing",
	},
	"LOF252": {
		ID: "LOF252",
		Name: "The Daughter",
	},
	"LOF253": {
		ID: "LOF253",
		Name: "Longbeam Cruiser",
	},
	"LOF254": {
		ID: "LOF254",
		Name: "Porg",
	},
	"LOF255": {
		ID: "LOF255",
		Name: "Curious Flock",
	},
	"LOF256": {
		ID: "LOF256",
		Name: "Gifted Urchin",
	},
	"LOF257": {
		ID: "LOF257",
		Name: "Kowakian Monkey-Lizard",
	},
	"LOF258": {
		ID: "LOF258",
		Name: "Peli Motto",
	},
	"LOF259": {
		ID: "LOF259",
		Name: "Ravening Gundark",
	},
	"LOF260": {
		ID: "LOF260",
		Name: "The Father",
	},
	"LOF261": {
		ID: "LOF261",
		Name: "Constructed Lightsaber",
	},
	"LOF262": {
		ID: "LOF262",
		Name: "Go Into Hiding",
	},
	"LOF263": {
		ID: "LOF263",
		Name: "Last Words",
	},
	"LOF264": {
		ID: "LOF264",
		Name: "It's Worse",
	},
}