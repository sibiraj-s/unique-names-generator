package unique

// Animals list
var Animals = []string{
	"canidae",
	"felidae",
	"cat",
	"cattle",
	"dog",
	"donkey",
	"goat",
	"horse",
	"pig",
	"rabbit",
	"aardvark",
	"aardwolf",
	"albatross",
	"alligator",
	"alpaca",
	"amphibian",
	"anaconda",
	"angelfish",
	"anglerfish",
	"ant",
	"anteater",
	"antelope",
	"antlion",
	"ape",
	"aphid",
	"armadillo",
	"asp",
	"baboon",
	"badger",
	"bandicoot",
	"barnacle",
	"barracuda",
	"basilisk",
	"bass",
	"bat",
	"bear",
	"beaver",
	"bedbug",
	"bee",
	"beetle",
	"bird",
	"bison",
	"blackbird",
	"boa",
	"boar",
	"bobcat",
	"bobolink",
	"bonobo",
	"booby",
	"bovid",
	"bug",
	"butterfly",
	"buzzard",
	"camel",
	"canid",
	"capybara",
	"cardinal",
	"caribou",
	"carp",
	"cat",
	"catshark",
	"caterpillar",
	"catfish",
	"cattle",
	"centipede",
	"cephalopod",
	"chameleon",
	"cheetah",
	"chickadee",
	"chicken",
	"chimpanzee",
	"chinchilla",
	"chipmunk",
	"clam",
	"clownfish",
	"cobra",
	"cockroach",
	"cod",
	"condor",
	"constrictor",
	"coral",
	"cougar",
	"cow",
	"coyote",
	"crab",
	"crane",
	"crawdad",
	"crayfish",
	"cricket",
	"crocodile",
	"crow",
	"cuckoo",
	"cicada",
	"damselfly",
	"deer",
	"dingo",
	"dinosaur",
	"dog",
	"dolphin",
	"donkey",
	"dormouse",
	"dove",
	"dragonfly",
	"dragon",
	"duck",
	"eagle",
	"earthworm",
	"earwig",
	"echidna",
	"eel",
	"egret",
	"elephant",
	"elk",
	"emu",
	"ermine",
	"falcon",
	"ferret",
	"finch",
	"firefly",
	"fish",
	"flamingo",
	"flea",
	"fly",
	"flyingfish",
	"fowl",
	"fox",
	"frog",
	"gamefowl",
	"galliform",
	"gazelle",
	"gecko",
	"gerbil",
	"gibbon",
	"giraffe",
	"goat",
	"goldfish",
	"goose",
	"gopher",
	"gorilla",
	"grasshopper",
	"grouse",
	"guan",
	"guanaco",
	"guineafowl",
	"gull",
	"guppy",
	"haddock",
	"halibut",
	"hamster",
	"hare",
	"harrier",
	"hawk",
	"hedgehog",
	"heron",
	"herring",
	"hippopotamus",
	"hookworm",
	"hornet",
	"horse",
	"hoverfly",
	"hummingbird",
	"hyena",
	"iguana",
	"impala",
	"jackal",
	"jaguar",
	"jay",
	"jellyfish",
	"junglefowl",
	"kangaroo",
	"kingfisher",
	"kite",
	"kiwi",
	"koala",
	"koi",
	"krill",
	"ladybug",
	"lamprey",
	"landfowl",
	"lark",
	"leech",
	"lemming",
	"lemur",
	"leopard",
	"leopon",
	"limpet",
	"lion",
	"lizard",
	"llama",
	"lobster",
	"locust",
	"loon",
	"louse",
	"lungfish",
	"lynx",
	"macaw",
	"mackerel",
	"magpie",
	"mammal",
	"manatee",
	"mandrill",
	"marlin",
	"marmoset",
	"marmot",
	"marsupial",
	"marten",
	"mastodon",
	"meadowlark",
	"meerkat",
	"mink",
	"minnow",
	"mite",
	"mockingbird",
	"mole",
	"mollusk",
	"mongoose",
	"monkey",
	"moose",
	"mosquito",
	"moth",
	"mouse",
	"mule",
	"muskox",
	"narwhal",
	"newt",
	"nightingale",
	"ocelot",
	"octopus",
	"opossum",
	"orangutan",
	"orca",
	"ostrich",
	"otter",
	"owl",
	"ox",
	"panda",
	"panther",
	"parakeet",
	"parrot",
	"parrotfish",
	"partridge",
	"peacock",
	"peafowl",
	"pelican",
	"penguin",
	"perch",
	"pheasant",
	"pig",
	"pigeon",
	"pike",
	"pinniped",
	"piranha",
	"planarian",
	"platypus",
	"pony",
	"porcupine",
	"porpoise",
	"possum",
	"prawn",
	"primate",
	"ptarmigan",
	"puffin",
	"puma",
	"python",
	"quail",
	"quelea",
	"quokka",
	"rabbit",
	"raccoon",
	"rat",
	"rattlesnake",
	"raven",
	"reindeer",
	"reptile",
	"rhinoceros",
	"roadrunner",
	"rodent",
	"rook",
	"rooster",
	"roundworm",
	"sailfish",
	"salamander",
	"salmon",
	"sawfish",
	"scallop",
	"scorpion",
	"seahorse",
	"shark",
	"sheep",
	"shrew",
	"shrimp",
	"silkworm",
	"silverfish",
	"skink",
	"skunk",
	"sloth",
	"slug",
	"smelt",
	"snail",
	"snake",
	"snipe",
	"sole",
	"sparrow",
	"spider",
	"spoonbill",
	"squid",
	"squirrel",
	"starfish",
	"stingray",
	"stoat",
	"stork",
	"sturgeon",
	"swallow",
	"swan",
	"swift",
	"swordfish",
	"swordtail",
	"tahr",
	"takin",
	"tapir",
	"tarantula",
	"tarsier",
	"termite",
	"tern",
	"thrush",
	"tick",
	"tiger",
	"tiglon",
	"toad",
	"tortoise",
	"toucan",
	"trout",
	"tuna",
	"turkey",
	"turtle",
	"tyrannosaurus",
	"urial",
	"vicuna",
	"viper",
	"vole",
	"vulture",
	"wallaby",
	"walrus",
	"wasp",
	"warbler",
	"weasel",
	"whale",
	"whippet",
	"whitefish",
	"wildcat",
	"wildebeest",
	"wildfowl",
	"wolf",
	"wolverine",
	"wombat",
	"woodpecker",
	"worm",
	"wren",
	"xerinae",
	"yak",
	"zebra",
	"alpaca",
	"cat",
	"cattle",
	"chicken",
	"dog",
	"donkey",
	"ferret",
	"gayal",
	"goldfish",
	"guppy",
	"horse",
	"koi",
	"llama",
	"sheep",
	"yak",
	"unicorn",
}
