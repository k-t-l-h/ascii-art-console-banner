package banner

// this is a fallback font
var defaultFont = "flf2a$ 8 8 17 -1 2\n"+
"basic.flf by Craig O'Flaherty <cofl@it.ntu.edu.au>\n"+
"August 17, 1994\n"+
"$$@\n"+
"$$@\n"+
"$$@\n"+
"$$@\n"+
"$$@\n"+
"$$@\n"+
"$$@\n"+
"$$@@\n"+
"db$@\n"+
"88$@\n"+
"YP$@\n"+
"  $@\n"+
"db$@\n"+
"YP$@\n"+
"  $@\n"+
"  $@@\n"+
".o. .o.$@\n"+
"`8' `8'$@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"       $@@\n"+
"       $@\n"+
" db db $@\n"+
"C88888D$@\n"+
" 88 88 $@\n"+
"C88888D$@\n"+
" YP YP $@\n"+
"       $@\n"+
"       $@@\n"+
"   A   $@\n"+
".d8888.$@\n"+
"88'8 YP$@\n"+
"`8b8.  $@\n"+
"  `V8b.$@\n"+
"db 8 8D$@\n"+
"`8888Y'$@\n"+
"   V   $@@\n"+
"db   dD$@\n"+
"YP  d8'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8' db$@\n"+
"d8'  YP$@\n"+
"       $@\n"+
"       $@@\n"+
".d888b. $@\n"+
"8P   8D $@\n"+
"`Vb d8' $@\n"+
" d88C dD$@\n"+
"C8' d8D $@\n"+
"`888P Yb$@\n"+
"        $@\n"+
"        $@@\n"+
"Cb$@\n"+
"`D$@\n"+
" '$@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@@\n"+
"    dD$@\n"+
"  d8' $@\n"+
" d8   $@\n"+
"C88   $@\n"+
" V8   $@\n"+
"  V8. $@\n"+
"    VD$@\n"+
"      $@@\n"+
"Cb    $@\n"+
" `8b  $@\n"+
"   8b $@\n"+
"   88D$@\n"+
"   8P $@\n"+
" .8P  $@\n"+
"CP    $@\n"+
"      $@@\n"+
"       $@\n"+
"8. A .8$@\n"+
"`8.8.8'$@\n"+
"  888  $@\n"+
".d'8`b.$@\n"+
"8' V `8$@\n"+
"       $@\n"+
"       $@@\n"+
"      $@\n"+
"  db  $@\n"+
"  88  $@\n"+
"C8888D$@\n"+
"  88  $@\n"+
"  VP  $@\n"+
"      $@\n"+
"      $@@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"db$@\n"+
"V8$@\n"+
" P$@\n"+
"  $@@\n"+
"      $@\n"+
"      $@\n"+
"      $@\n"+
"C8888D$@\n"+
"      $@\n"+
"      $@\n"+
"      $@\n"+
"      $@@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"db$@\n"+
"VP$@\n"+
"  $@\n"+
"  $@@\n"+
"     dD$@\n"+
"    d8'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8'   $@\n"+
"C8'    $@\n"+
"       $@\n"+
"       $@@\n"+
" .d88b. $@\n"+
".8P  88.$@\n"+
"88  d'88$@\n"+
"88 d' 88$@\n"+
"`88  d8'$@\n"+
" `Y88P' $@\n"+
"        $@\n"+
"        $@@\n"+
" db$@\n"+
"o88$@\n"+
" 88$@\n"+
" 88$@\n"+
" 88$@\n"+
" VP$@\n"+
"   $@\n"+
"   $@@\n"+
".d888b.$@\n"+
"VP  `8D$@\n"+
"   odD'$@\n"+
" .88'  $@\n"+
"j88.   $@\n"+
"888888D$@\n"+
"       $@\n"+
"       $@@\n"+
"d8888b.$@\n"+
"VP  `8D$@\n"+
"  oooY'$@\n"+
"  ~~~b.$@\n"+
"db   8D$@\n"+
"Y8888P'$@\n"+
"       $@\n"+
"       $@@\n"+
"  j88D $@\n"+
" j8~88 $@\n"+
"j8' 88 $@\n"+
"V88888D$@\n"+
"    88 $@\n"+
"    VP $@\n"+
"       $@\n"+
"       $@@\n"+
"  ooooo$@\n"+
" 8P~~~~$@\n"+
"dP     $@\n"+
"V8888b.$@\n"+
"    `8D$@\n"+
"88oobY'$@\n"+
"       $@\n"+
"       $@@\n"+
"   dD  $@\n"+
"  d8'  $@\n"+
" d8'   $@\n"+
"d8888b.$@\n"+
"88' `8D$@\n"+
"`8888P $@\n"+
"       $@\n"+
"       $@@\n"+
"d88888D$@\n"+
"VP  d8'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8'   $@\n"+
"d8'    $@\n"+
"       $@\n"+
"       $@@\n"+
".d888b.$@\n"+
"88   8D$@\n"+
"`VoooY'$@\n"+
".d~~~b.$@\n"+
"88   8D$@\n"+
"`Y888P'$@\n"+
"       $@\n"+
"       $@@\n"+
".d888b.$@\n"+
"88' `8D$@\n"+
"`V8o88'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8'   $@\n"+
"       $@\n"+
"       $@@\n"+
"  $@\n"+
"db$@\n"+
"VP$@\n"+
"  $@\n"+
"db$@\n"+
"VP$@\n"+
"  $@\n"+
"  $@@\n"+
"  $@\n"+
"db$@\n"+
"VP$@\n"+
"  $@\n"+
"db$@\n"+
"V8$@\n"+
" P$@\n"+
"  $@@\n"+
"      $@\n"+
"   .dP$@\n"+
" .d8  $@\n"+
",P    $@\n"+
"`b    $@\n"+
" `Vb  $@\n"+
"   `Vb$@\n"+
"      $@@\n"+
"      $@\n"+
"C8888D$@\n"+
"      $@\n"+
"C8888D$@\n"+
"      $@\n"+
"      $@\n"+
"      $@\n"+
"      $@@\n"+
"      $@\n"+
"Vb    $@\n"+
" `Vb  $@\n"+
"   `V.$@\n"+
"   .d'$@\n"+
" .dP  $@\n"+
"dP    $@\n"+
"      $@@\n"+
".d888b.$@\n"+
"VP  `8D$@\n"+
"   odD'$@\n"+
"  8P'  $@\n"+
"  oo   $@\n"+
"  VP   $@\n"+
"       $@\n"+
"       $@@\n"+
" .o888b.$@\n"+
"d8'   Y8$@\n"+
"8P db dP$@\n"+
"8b V8o8P$@\n"+
"Y8.    d$@\n"+
" `Y888P'$@\n"+
"        $@\n"+
"        $@@\n"+
" .d8b. $@\n"+
"d8' `8b$@\n"+
"88ooo88$@\n"+
"88~~~88$@\n"+
"88   88$@\n"+
"YP   YP$@\n"+
"       $@\n"+
"       $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oooY'$@\n"+
"88~~~b.$@\n"+
"88   8D$@\n"+
"Y8888P'$@\n"+
"       $@\n"+
"       $@@\n"+
" .o88b.$@\n"+
"d8P  Y8$@\n"+
"8P     $@\n"+
"8b     $@\n"+
"Y8b  d8$@\n"+
" `Y88P'$@\n"+
"       $@\n"+
"       $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88   88$@\n"+
"88   88$@\n"+
"88  .8D$@\n"+
"Y8888D'$@\n"+
"       $@\n"+
"       $@@\n"+
"d88888b$@\n"+
"88'    $@\n"+
"88ooooo$@\n"+
"88~~~~~$@\n"+
"88.    $@\n"+
"Y88888P$@\n"+
"       $@\n"+
"       $@@\n"+
"d88888b$@\n"+
"88'    $@\n"+
"88ooo  $@\n"+
"88~~~  $@\n"+
"88     $@\n"+
"YP     $@\n"+
"       $@\n"+
"       $@@\n"+
" d888b $@\n"+
"88' Y8b$@\n"+
"88     $@\n"+
"88  ooo$@\n"+
"88. ~8~$@\n"+
" Y888P $@\n"+
"       $@\n"+
"       $@@\n"+
"db   db$@\n"+
"88   88$@\n"+
"88ooo88$@\n"+
"88~~~88$@\n"+
"88   88$@\n"+
"YP   YP$@\n"+
"       $@\n"+
"       $@@\n"+
"d888888b$@\n"+
"  `88'  $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"  .88.  $@\n"+
"Y888888P$@\n"+
"        $@\n"+
"        $@@\n"+
"   d88b$@\n"+
"   `8P'$@\n"+
"    88 $@\n"+
"    88 $@\n"+
"db. 88 $@\n"+
"Y8888P $@\n"+
"       $@\n"+
"       $@@\n"+
"db   dD$@\n"+
"88 ,8P'$@\n"+
"88,8P  $@\n"+
"88`8b  $@\n"+
"88 `88.$@\n"+
"YP   YD$@\n"+
"       $@\n"+
"       $@@\n"+
"db     $@\n"+
"88     $@\n"+
"88     $@\n"+
"88     $@\n"+
"88booo.$@\n"+
"Y88888P$@\n"+
"       $@\n"+
"       $@@\n"+
".88b  d88.$@\n"+
"88'YbdP`88$@\n"+
"88  88  88$@\n"+
"88  88  88$@\n"+
"88  88  88$@\n"+
"YP  YP  YP$@\n"+
"          $@\n"+
"          $@@\n"+
"d8b   db$@\n"+
"888o  88$@\n"+
"88V8o 88$@\n"+
"88 V8o88$@\n"+
"88  V888$@\n"+
"VP   V8P$@\n"+
"        $@\n"+
"        $@@\n"+
" .d88b. $@\n"+
".8P  Y8.$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"`8b  d8'$@\n"+
" `Y88P' $@\n"+
"        $@\n"+
"        $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oodD'$@\n"+
"88~~~  $@\n"+
"88     $@\n"+
"88     $@\n"+
"       $@\n"+
"       $@@\n"+
" .d88b. $@\n"+
".8P  Y8.$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"`8P  d8'$@\n"+
" `Y88'Y8$@\n"+
"        $@\n"+
"        $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oobY'$@\n"+
"88`8b  $@\n"+
"88 `88.$@\n"+
"88   YD$@\n"+
"       $@\n"+
"       $@@\n"+
".d8888.$@\n"+
"88'  YP$@\n"+
"`8bo.  $@\n"+
"  `Y8b.$@\n"+
"db   8D$@\n"+
"`8888Y'$@\n"+
"       $@\n"+
"       $@@\n"+
"d888888b$@\n"+
"`~~88~~'$@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"88b  d88$@\n"+
"~Y8888P'$@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"88    88$@\n"+
"Y8    8P$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"db   d8b   db$@\n"+
"88   I8I   88$@\n"+
"88   I8I   88$@\n"+
"Y8   I8I   88$@\n"+
"`8b d8'8b d8'$@\n"+
" `8b8' `8d8' $@\n"+
"             $@\n"+
"             $@@\n"+
"db    db$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
" .dPYb. $@\n"+
".8P  Y8.$@\n"+
"YP    YP$@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"d88888D$@\n"+
"YP  d8'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8' db$@\n"+
"d88888P$@\n"+
"       $@\n"+
"       $@@\n"+
"d88D$@\n"+
"88  $@\n"+
"88  $@\n"+
"88  $@\n"+
"88  $@\n"+
"L88D$@\n"+
"    $@\n"+
"    $@@\n"+
"Cb     $@\n"+
"`8b    $@\n"+
" `8b   $@\n"+
"  `8b  $@\n"+
"   `8b $@\n"+
"    `8D$@\n"+
"       $@\n"+
"       $@@\n"+
"C88D$@\n"+
"  88$@\n"+
"  88$@\n"+
"  88$@\n"+
"  88$@\n"+
"C888$@\n"+
"    $@\n"+
"    $@@\n"+
"   db   $@\n"+
" .dPVb. $@\n"+
"dP'  `Vb$@\n"+
"        $@\n"+
"        $@\n"+
"        $@\n"+
"        $@\n"+
"        $@@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"       $@\n"+
"C88888D$@\n"+
"       $@\n"+
"       $@@\n"+
"dD$@\n"+
"C'$@\n"+
" `$@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@\n"+
"  $@@\n"+
" .d8b. $@\n"+
"d8' `8b$@\n"+
"88ooo88$@\n"+
"88~~~88$@\n"+
"88   88$@\n"+
"YP   YP$@\n"+
"       $@\n"+
"       $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oooY'$@\n"+
"88~~~b.$@\n"+
"88   8D$@\n"+
"Y8888P'$@\n"+
"       $@\n"+
"       $@@\n"+
" .o88b.$@\n"+
"d8P  Y8$@\n"+
"8P     $@\n"+
"8b     $@\n"+
"Y8b  d8$@\n"+
" `Y88P'$@\n"+
"       $@\n"+
"       $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88   88$@\n"+
"88   88$@\n"+
"88  .8D$@\n"+
"Y8888D'$@\n"+
"       $@\n"+
"       $@@\n"+
"d88888b$@\n"+
"88'    $@\n"+
"88ooooo$@\n"+
"88~~~~~$@\n"+
"88.    $@\n"+
"Y88888P$@\n"+
"       $@\n"+
"       $@@\n"+
"d88888b$@\n"+
"88'    $@\n"+
"88ooo  $@\n"+
"88~~~  $@\n"+
"88     $@\n"+
"YP     $@\n"+
"       $@\n"+
"       $@@\n"+
" d888b $@\n"+
"88' Y8b$@\n"+
"88     $@\n"+
"88  ooo$@\n"+
"88. ~8~$@\n"+
" Y888P $@\n"+
"       $@\n"+
"       $@@\n"+
"db   db$@\n"+
"88   88$@\n"+
"88ooo88$@\n"+
"88~~~88$@\n"+
"88   88$@\n"+
"YP   YP$@\n"+
"       $@\n"+
"       $@@\n"+
"d888888b$@\n"+
"  `88'  $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"  .88.  $@\n"+
"Y888888P$@\n"+
"        $@\n"+
"        $@@\n"+
"   d88b$@\n"+
"   `8P'$@\n"+
"    88 $@\n"+
"    88 $@\n"+
"db. 88 $@\n"+
"Y8888P $@\n"+
"       $@\n"+
"       $@@\n"+
"db   dD$@\n"+
"88 ,8P'$@\n"+
"88,8P  $@\n"+
"88`8b  $@\n"+
"88 `88.$@\n"+
"YP   YD$@\n"+
"       $@\n"+
"       $@@\n"+
"db     $@\n"+
"88     $@\n"+
"88     $@\n"+
"88     $@\n"+
"88booo.$@\n"+
"Y88888P$@\n"+
"       $@\n"+
"       $@@\n"+
".88b  d88.$@\n"+
"88'YbdP`88$@\n"+
"88  88  88$@\n"+
"88  88  88$@\n"+
"88  88  88$@\n"+
"YP  YP  YP$@\n"+
"          $@\n"+
"          $@@\n"+
"d8b   db$@\n"+
"888o  88$@\n"+
"88V8o 88$@\n"+
"88 V8o88$@\n"+
"88  V888$@\n"+
"VP   V8P$@\n"+
"        $@\n"+
"        $@@\n"+
" .d88b. $@\n"+
".8P  Y8.$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"`8b  d8'$@\n"+
" `Y88P' $@\n"+
"        $@\n"+
"        $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oodD'$@\n"+
"88~~~  $@\n"+
"88     $@\n"+
"88     $@\n"+
"       $@\n"+
"       $@@\n"+
" .d88b. $@\n"+
".8P  Y8.$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"`8P  d8'$@\n"+
" `Y88'Y8$@\n"+
"        $@\n"+
"        $@@\n"+
"d8888b.$@\n"+
"88  `8D$@\n"+
"88oobY'$@\n"+
"88`8b  $@\n"+
"88 `88.$@\n"+
"88   YD$@\n"+
"       $@\n"+
"       $@@\n"+
".d8888.$@\n"+
"88'  YP$@\n"+
"`8bo.  $@\n"+
"  `Y8b.$@\n"+
"db   8D$@\n"+
"`8888Y'$@\n"+
"       $@\n"+
"       $@@\n"+
"d888888b$@\n"+
"`~~88~~'$@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"88    88$@\n"+
"88b  d88$@\n"+
"~Y8888P'$@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"88    88$@\n"+
"Y8    8P$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"db   d8b   db$@\n"+
"88   I8I   88$@\n"+
"88   I8I   88$@\n"+
"Y8   I8I   88$@\n"+
"`8b d8'8b d8'$@\n"+
" `8b8' `8d8' $@\n"+
"             $@\n"+
"             $@@\n"+
"db    db$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
" .dPYb. $@\n"+
".8P  Y8.$@\n"+
"YP    YP$@\n"+
"        $@\n"+
"        $@@\n"+
"db    db$@\n"+
"`8b  d8'$@\n"+
" `8bd8' $@\n"+
"   88   $@\n"+
"   88   $@\n"+
"   YP   $@\n"+
"        $@\n"+
"        $@@\n"+
"d88888D$@\n"+
"YP  d8'$@\n"+
"   d8' $@\n"+
"  d8'  $@\n"+
" d8' db$@\n"+
"d88888P$@\n"+
"       $@\n"+
"       $@@\n"+
"   .8P$@\n"+
"   8' $@\n"+
" .dP  $@\n"+
"C88   $@\n"+
" `Yb  $@\n"+
"   8. $@\n"+
"   `8b$@\n"+
"      $@@\n"+
"8$@\n"+
"8$@\n"+
"8$@\n"+
" $@\n"+
"8$@\n"+
"8$@\n"+
"8$@\n"+
" $@@\n"+
"V8.   $@\n"+
" `8   $@\n"+
"  Vb. $@\n"+
"   88D$@\n"+
"  dP' $@\n"+
" .8   $@\n"+
"C8'   $@\n"+
"      $@@\n"+
" .oo.  .$@\n"+
"P'  `VP'$@\n"+
"        $@\n"+
"        $@\n"+
"        $@\n"+
"        $@\n"+
"        $@\n"+
"        $@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@\n"+
"@@\n"+
""