module TwelveDays exposing (countEnglish, recite)

import Dict


recite : Int -> Int -> List String
recite _ _ =
    [ "Please implement this function" ]


countEnglish : Int -> String
countEnglish n =
    let
        d =
            Dict.fromList
                [ ( 1, "first" )
                , ( 2, "second" )
                ]
    in
    case Dict.get n d of
        Nothing ->
            ""

        Just word ->
            word
