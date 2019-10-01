module TwelveDays exposing (countEnglish, gifts, recite)

import Dict
import List
import Maybe
import String


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
    Maybe.withDefault "" (Dict.get n d)


gifts n =
    let
        l =
            [ "two Turtle Doves"
            , "three French Hens"
            , "four Calling Birds"
            , "five Gold Rings"
            , "six Geese-a-Laying"
            , "seven Swans-a-Swimming"
            , "eight Maids-a-Milking"
            , "nine Ladies Dancing"
            , "ten Lords-a-Leaping"
            , "eleven Pipers Piping"
            , "twelve Drummers Drumming"
            ]
    in
    if n == 1 then
        "a Partridge in a Pear Tree"

    else
        String.join ", " (List.reverse (List.take (n - 1) l)) ++ ", and a Partridge in a Pear Tree"
