package games

var Hangman_display [7]string = [7]string{
	`
  +---+
 |      |
        |
        |
        |
        |
=======`,
`
	+---+
     |      |
    O     |
		    |
		    |
		    |
    =======`,
 `
 +---+
  |      |
 O     |
  |      |
		 |
		 |
 =======`,
`
 +---+
  |      |
 O     |
/|      |
		 |
		 |
  =======`,
 `
 +---+
  |      |
 O     |
/|\    |
		 |
		 |
   =======`,
	 `
	 +---+
	  |      |
	 O     |
	/|\    |
	/       |
			 |
   =======`,
 `
 +---+
  |      |
 O     |
/|\    |
/ \    |
		  |
  =======`}