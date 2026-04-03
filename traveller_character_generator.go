type Stat struct {
    strengh         int range 1 to 15
    dexterity       int range 1 to 15
    endurance       int range 1 to 15
    intelligence    int range 1 to 15
    education       int range 1 to 15
    socialStatus    int range 1 to 15
}

type Character struct {
    name            string
    age             int = 18
    stat            Stat
    title           string
    rank            string
    service         string
    termsServed     int = 0
    isRetired       bool = false
    retirementPay   int
    isTASMember     bool = false
}

type Career struct {
    name        string
    enlistment  int
    draft       int
    survival    int
    commission  int
    promotion   int
    reenlist    int
}