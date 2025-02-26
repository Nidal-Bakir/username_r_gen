package usernaemgen

type DictType byte

const (
	// Adjective word list.
	Adjectives DictType = iota
	// Color word list.
	Colors DictType = iota
	// Animal word list.
	Animals DictType = iota
	// Max value of DictType. Used internally for bounds checking.
	numDicts DictType = iota
)

var dicts [numDicts][]string
var DictsLen [numDicts]int

func init() {
	dicts[Adjectives] = adjectives
	dicts[Colors] = colors
	dicts[Animals] = animals

	DictsLen[Adjectives] = len(adjectives)
	DictsLen[Colors] = len(colors)
	DictsLen[Animals] = len(animals)
}

var adjectives = []string{
	"abbreviated", "cute", "aberrant", "abhorrent", "abiding", "abnormal", "abundant", "adorable", "adventurous", "aggressive",
	"ambitious", "ancient", "angry", "animated", "annoying", "antsy", "anxious", "arrogant", "articulate", "ashamed",
	"attractive", "auspicious", "awesome", "awful", "bald", "barren", "baseless", "bashful", "basic", "beautiful",
	"best", "big", "bitter", "bizarre", "blissful", "bored", "boring", "boujee", "brainy", "bright",
	"broad", "broken", "busy", "calm", "capable", "careful", "careless", "caring", "cautious", "charming",
	"chatty", "cheerful", "chubby", "clean", "clever", "clumsy", "cold", "colorful", "comfortable", "concerned",
	"confused", "curious", "curly", "curteous", "curvy", "cute", "daft", "dainty", "damaged", "damp",
	"dangerous", "dark", "deadly", "deep", "defective", "delicate", "determined", "different", "dirty", "docile",
	"dry", "dusty", "early", "educated", "efficient", "elderly", "elegant", "embarrassed", "empty", "encouraging",
	"enthusiastic", "excellent", "exciting", "expensive", "fabulous", "fair", "faithful", "famous", "fancy", "fantastic",
	"fast", "fearful", "fearless", "fertile", "flamboyant", "foolish", "forgetful", "friendly", "funny", "gentle",
	"glamorous", "glorious", "goofy", "gorgeous", "graceful", "grateful", "great", "greedy", "hairy", "handsome",
	"happy", "harsh", "hasty", "healthy", "heavy", "helpful", "hilarious", "historical", "hostile", "hot",
	"huge", "humorous", "hungry", "hyperactive", "hyperbolic", "hysterical", "ignorant", "illegal", "imaginary", "impolite",
	"important", "impossible", "innocent", "inspiring", "intelligent", "interesting", "jealous", "jolly", "juicy", "jumpy",
	"juvenile", "killer", "kind", "large", "lazy", "legal", "light", "literate", "little", "lively",
	"livid", "lonely", "loopy", "loud", "lovely", "lucky", "macho", "magical", "magnificent", "massive",
	"mature", "mean", "mega", "messy", "mighty", "miniature", "modern", "muscular", "narrow", "nasty",
	"naughty", "nervous", "new", "noisy", "nutritious", "obedient", "obese", "old", "overconfident", "peaceful",
	"pious", "polite", "poor", "powerful", "precious", "pretty", "proud", "quick", "quiet", "rapid",
	"rare", "remarkable", "responsible", "rich", "romantic", "royal", "rude", "scintillating", "secretive", "selfish",
	"serious", "sharp", "shiny", "shocking", "short", "shy", "silly", "sincere", "skinny", "slim",
	"slow", "small", "snooty", "soft", "space", "spicy", "spiritual", "splendid", "strong", "successful",
	"supercharged", "sweet", "tactful", "talented", "tall", "tangible", "tasteful", "tasty", "temperate", "tenacious",
	"tender", "tense", "terrific", "thankful", "thick", "thin", "thorough", "thoughtful", "tiny", "unique",
	"untidy", "upbeat", "upset", "victorious", "violent", "warm", "weak", "wealthy", "weary", "wide", "alien", "aloof",
	"wise", "witty", "wonderful", "worried", "young", "youthful", "zealous", "zen", "Elegant", "Gorgeous",
	"Stunning", "Radiant", "Majestic", "Picturesque", "Ravishing", "Amiable", "Charismatic", "Compassionate", "Diligent", "Generous",
	"Gentle", "Humble", "Optimistic", "Reliable", "Resilient", "Brilliant", "Ingenious", "Innovative", "Logical", "Perceptive",
	"Quick-witted", "Resourceful", "Accomplished", "Admirable", "Confident", "Determined", "Fearless", "Industrious", "Prolific", "Triumphant",
	"Colossal", "Compact", "Elongated", "Gargantuan", "Minuscule", "Rotund", "Slender", "Towering", "Brisk", "Hasty",
	"Lethargic", "Meandering", "Sluggish", "Sprightly", "Unhurried", "Zippy", "Brittle", "Coarse", "Fluffy", "Gritty",
	"Lush", "Sleek", "Squishy", "Velvety",
}

var colors = []string{
	"blue", "beige", "black", "aqua", "bronze", "brown", "burgundy", "charcoal", "chocolate", "cyan",
	"emerald", "fuchsia", "gold", "gray", "green", "ivory", "jade", "lavender", "lime", "magenta",
	"maroon", "mint", "mustard", "navy", "ochre", "olive", "orange", "orchid", "pale", "peach",
	"periwinkle", "pink", "plum", "purple", "red", "rose", "ruby", "saffron", "silver", "slate",
	"tan", "teal", "turquoise", "violet", "white", "yellow", "crimson", "Red", "Blue", "Green",
	"Yellow", "Orange", "Purple", "Pink", "Brown", "Black", "White", "Gray", "Cyan", "Magenta",
	"Lime", "Olive", "Teal", "Maroon", "Navy", "Beige", "Turquoise", "Gold", "Silver", "Bronze",
	"Lavender", "Indigo", "Coral", "Peach", "Mint", "Salmon", "Crimson", "Periwinkle", "Charcoal", "Plum",
	"Azure", "Vermilion", "Cerulean", "Fuchsia", "Emerald", "Jade", "Ruby", "Saffron", "Ochre", "Amethyst",
	"Topaz", "Ivory", "Mustard", "Garnet", "Copper", "Rose", "Eggplant", "Lilac", "Mahogany", "Obsidian",
	"Cobalt", "Burgundy", "Sapphire", "Zaffre", "Amber", "Quartz", "Sand", "Sepia", "Slate", "Tangerine",
	"Verdigris", "Chartreuse", "Moss", "Carmine", "Bistre", "Onyx", "Vanilla", "Heliotrope", "Mulberry", "Celadon",
	"Auburn", "Taupe", "Denim", "Brass", "Russet", "Pewter", "Wisteria", "Fern", "Mauve", "Ultramarine",
	"Titanium", "Ash", "Opal", "Ebony", "Persimmon", "Basil", "Blush", "Grape", "Cerise", "Teaberry",
	"Storm", "Pine", "Seafoam", "Gunmetal", "Hickory", "Snow", "Fog", "amber",
}

var animals = []string{
	"fox", "albatross", "alligator", "alpaca", "anchovy", "angler", "ant", "anteater", "antelope", "ape", "armadillo",
	"arowana", "asp", "auk", "avocet", "ayeaye", "baboon", "badger", "bandicoot", "banteng", "barnacle",
	"barracuda", "barramundi", "basilisk", "bass", "bat", "beagle", "bear", "beaver", "bedbug", "bee",
	"beetle", "bichon", "bird", "bison", "bittern", "blackbird", "blobfish", "bloodhound", "blowfish", "bluebird",
	"bluejay", "boa", "bobcat", "bonito", "bonobo", "booby", "bovidae", "bream", "brent", "brill",
	"brolga", "bronco", "brownie", "buffalo", "bulldog", "bullfrog", "bumblebee", "bunting", "butterfly", "buzzard",
	"caiman", "calf", "camel", "canary", "capuchin", "capybara", "caracal", "cardinal", "caribou", "carp",
	"cat", "catbird", "caterpillar", "catfish", "cattle", "centipede", "chameleon", "chamois", "cheetah", "chicken",
	"chihuahua", "chimpanzee", "chinchilla", "chipmunk", "chough", "chowchow", "cicadia", "civet", "clam", "clamworm",
	"clownfish", "cobra", "cockatoo", "cockroach", "cod", "coelacanth", "collie", "colobus", "condor", "conure",
	"coorgi", "coral", "cormorant", "cougar", "cow", "coyote", "crab", "crane", "crappie", "crayfish",
	"creeper", "cricket", "croc", "crocodile", "crow", "cuckoo", "curlew", "cuttlefish", "dace", "dachshund",
	"dalmatian", "damselfly", "dassie", "deer", "dhole", "dingo", "dinosaur", "dipper", "discus", "doberman",
	"dodo", "dog", "dogfish", "doggo", "doggy", "dolphin", "dormouse", "dove", "dragon",
	"dragonfly", "drake", "dromedary", "duck", "dugong", "dunlin", "eagle", "earwig", "echidna", "eel",
	"eland", "elephant", "elephantseal", "elk", "emu", "ermine", "eskimo", "falcon", "fennec", "ferret",
	"finch", "firefly", "fish", "flamingo", "flea", "flounder", "fly", "flyingfish", "fossa", "fowl",
	"frog", "fulmar", "gadwall", "galago", "gallinule", "gannet", "gar", "gazelle", "gecko",
	"gerbil", "gharial", "gibbon", "gila", "giraffe", "glider", "glowworm", "gnat", "gnu", "goat", "aardvark",
	"godwit", "goldeneye", "goldfinch", "goldfish", "goose", "gopher", "gorilla", "goshawk", "gourami", "grackle",
	"grasshopper", "grebe", "greyhound", "grizzly", "groundhog", "grouper", "grouse", "guanaco", "gudgeon", "guineafowl",
	"guineapig", "guppy", "gurnard", "gypsy", "gyrfalcon", "hammerhead", "hamster", "hare", "harpy", "harrier",
	"hawk", "hedgehog", "heron", "herring", "hippo", "hippopotamus", "hoatzin", "hog", "hoiho", "hoki",
	"honeybee", "hornbill", "hornet", "horse", "hoverfly", "howler", "human", "hummingbird", "husky", "hyena",
	"ibex", "ibis", "iguana", "impala", "indri", "jackal", "jackrabbit", "jaguar", "javelina", "jay",
	"jellyfish", "joey", "junco", "kakapo", "kangaroo", "katydid", "kingbird", "kingfisher", "kinglet", "kinkajou",
	"kite", "kiwi", "koala", "koi", "komodo", "kookaburra", "kouprey", "krill", "kudu", "lab", "zorilla",
	"labrador", "ladybug", "lamprey", "langur", "lark", "leech", "lemming", "lemur", "leopard", "liger",
	"limpet", "ling", "lion", "lionfish", "lizard", "llama", "lobster", "locust", "loon", "loris",
	"louse", "lovebird", "lynx", "macaw", "mackenzie", "mackerel", "magpie", "mako", "mallard", "mamba",
	"manatee", "mandrill", "manta", "mantaray", "mantis", "marmoset", "marmot", "marsupial", "martin", "mastiff",
	"mayfly", "meadowlark", "meerkat", "megalodon", "mink", "mite", "mola", "mole", "molly", "mongoose",
	"mongrel", "monkey", "moose", "mosquito", "moth", "mouse", "mule", "muntjac", "murre", "muskrat",
	"mussel", "mustang", "myotis", "nandu", "narwhal", "natterjack", "nautilus", "neanderthal", "newfoundland", "newt",
	"nighthawk", "nightingale", "numbat", "nutcracker", "nuthatch", "ocelot", "octopus", "okapi", "olm", "opossum",
	"orangutan", "orca", "oriole", "oryx", "osprey", "ostrich", "otter", "owl", "ox", "oxpecker", "wildebeest",
	"oyster", "pademelon", "panda", "pangolin", "panther", "parakeet", "parrot", "partridge", "peacock", "peafowl",
	"peccary", "pekingese", "pelican", "penguin", "perch", "petrel", "pheasant", "pigeon", "pika",
	"pikachu", "pike", "pilotfish", "pipit", "piranha", "pitbull", "platypus", "plover", "polarbear", "pollock",
	"polyp", "poodle", "porcupine", "porpoise", "possum", "prairie", "prawn", "pufferfish", "puffin", "pug",
	"puma", "pup", "puppy", "python", "quail", "quelea", "quetzal", "quokka", "quoll", "rabbit", "viper",
	"raccoon", "ragfish", "rail", "ram", "rat", "rattail", "rattlesnake", "raven", "razorbill", "reindeer",
	"reptile", "retriever", "rhea", "rhino", "rhinoceros", "roadrunner", "robin", "rook", "rooster", "rottweiler",
	"sable", "sailfish", "salamander", "salmon", "sanddollar", "sandpiper", "sapsucker", "sardine", "sawfish", "scallop",
	"scarab", "scorpion", "scoter", "seabass", "seagull", "seahorse", "seal", "sealion", "serval", "shad",
	"shark", "sheep", "sheepdog", "shibainu", "shihtzu", "shrew", "shrike", "shrimp", "silkworm", "silverfish",
	"skimmer", "skink", "skipper", "skua", "skunk", "sloth", "slug", "smelt", "snail", "snake", "wallaby",
	"snipe", "snowgoose", "snowshoe", "snowyowl", "sockeye", "solenodon", "sparrow", "sparrowhawk", "spider", "spiketail",
	"spittlebug", "sponge", "spoonbill", "springbok", "squid", "squirrel", "starfish", "starling", "steelhead", "stickleback",
	"stingray", "stoat", "stonefly", "stork", "studfish", "sugar", "sunbittern", "sunfish", "surgeonfish", "swan",
	"swift", "terrier", "tiger", "tortoise", "treefrog", "tuna", "turkey", "turtle", "unicorn", "vulture",
	"walrus", "wasp", "weasel", "whale", "wildboar", "wolf", "wolverine", "wombat", "woodpecker", "worm",
	"yak", "yeti", "zebra", "cicada", "corgi", "tamarin", "tapir", "tarsier", "termite", "toad", "toucan", "xerus",
}
