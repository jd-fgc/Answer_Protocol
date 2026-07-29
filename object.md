1. COMMUNICATION RÉSEAU

Message
  |-- From       string        (adresse conn.RemoteAddr(), identifiant
  |                              stable de la connexion, avant même CONNECT)
  |-- Text       string        (la ligne brute envoyée par le client)
  '-- ReplyChan  chan string   (channel privé pour renvoyer LA réponse
                                 à CE client précis)

  Rôle : c'est l'enveloppe que chaque handleConnection() envoie
  au Hub via le channel "incoming". Le Hub répond en écrivant
  dans ReplyChan.


2. SESSION

Session
  '-- Player  *Player   (nil tant que CONNECT n'a pas réussi)
  |-- username string 

  var sessions = make(map[string]*Session)
  clé de la map = Message.From (adresse de connexion)

  Rôle : fait le lien entre "une connexion TCP" (identifiée par
  son adresse) et "un joueur logique" (Player), une fois que
  CONNECT <nom> a été validé par le Hub.


3. JOUEURS / PERSONNAGES (models.go)

Players
  '-- Players  []Player

  Deux usages différents dans le code actuel :
    - playable_characters : personnages jouables,
      chargé depuis Characters.json
    - online_players : liste des joueurs actuellement connectés
      en partie

Player
  |-- Name          string
  |-- Item          string
  |-- Inventory     []string
  |-- MAX_HP        int
  |-- CURRENT_HP    int
  |-- ATK           float32
  |-- DMG_Mult      float32
  |-- Luck          int
  |-- Deal_chance   int
  |-- Respawn       int
  '-- Zone          string


4. MONDE / CARTE (models.go)

Rooms
  '-- Rooms  []Room   (json:"world" — chargé depuis world_map.json)

Room
  |-- Name    string
  |-- NPC     []NPC
  |-- Items   []string
  |-- North   string
  |-- South   string
  |-- East    string
  '-- West    string

NPC
  |-- Name     string
  |-- Talk     string   (dialogue renvoyée par TALK)
  |-- HP       int
  '-- Ennemy   bool


5. AFFICHAGE / RÉPONSE JSON POUR LOOK (look.go)

Look_Struct
  |-- Room     Room_Description
  |-- Players  []string
  |-- Items    []string
  '-- NPCs     []string

Room_Description
  |-- Id            string
  |-- Name          string
  |-- Description   string
  '-- Exits         Exits

Exits
  |-- North   string   (json:"north,omitempty")
  |-- South   string   (json:"south,omitempty")
  |-- East    string   (json:"east,omitempty")
  '-- West    string   (json:"west,omitempty")


RELATIONS

  Message.From ──────────────► clé dans sessions
                                       │
                                       ▼
                                Session.Player ──────► *Player
                                                           │
                                                           │ .Zone
                                                           ▼
  Rooms.Rooms[] ◄──────────────────────────────────── Room.Name
        │
        ├── Room.NPC[]     ──► NPC (Name, Talk, HP, Ennemy)
        ├── Room.Items[]   ──► []string
        └── Room.North/South/East/West ──► noms d'autres Room

  CONNECT usr_name
        │
        ├── Server return OK si la conn est etablie
        └── Envoi la selection des persos

      QUIT
        │
        ├── Enleve le joueur de la session ? 
        └── Server return un OK Disconnected ?


  CHOOSE character
        │
        └── Server return character=character_name


  LOOK(player, zones, players) ──► construit un Look_Struct
                                    (via Room_Description + Exits)
                                    ──► json.Marshal ──► envoyé au client


FICHIERS DE DONNÉES ↔ STRUCTS

  data/Characters.json   ──►  []Player
  data/world_map.json    ──►  Rooms
  data/quests.json       ──►  in progress