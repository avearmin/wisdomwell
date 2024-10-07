module HomePage exposing (main)

import Browser
import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Http exposing (expectJson, get)
import Json.Decode as Decode


type alias Quote =
    { id : String
    , createdAt : String
    , updatedAt : String
    , authorId : String
    , authorName : String
    , content : String
    }


type alias Model =
    { apiUrl : String
    , quote : Maybe Quote
    , isLoading : Bool
    , error : Maybe String
    }


type Msg
    = FetchQuote
    | FetchQuoteResponse (Result Http.Error Quote)


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        FetchQuote ->
            ( { model | isLoading = True, error = Nothing }
            , fetchQuoteFromBackend model.apiUrl
            )

        _ ->
            ( model
            , Cmd.none
            )


quoteDecoder : Decode.Decoder Quote
quoteDecoder =
    Decode.map6 Quote
        (Decode.field "id" Decode.string)
        (Decode.field "createdAt" Decode.string)
        (Decode.field "updatedAt" Decode.string)
        (Decode.field "authorId" Decode.string)
        (Decode.field "authorName" Decode.string)
        (Decode.field "content" Decode.string)


fetchQuoteFromBackend : String -> Cmd Msg
fetchQuoteFromBackend apiUrl =
    Http.get
        { url = apiUrl ++ "/quotes/random"
        , expect = Http.expectJson FetchQuoteResponse quoteDecoder
        }


init : String -> Model
init apiUrl =
    { apiUrl = apiUrl
    , quote = Nothing
    , isLoading = False
    , error = Nothing
    }


view : Model -> Html Msg
view model =
    div []
        [ button [ onClick FetchQuote ] [ text "Get Random Quote" ]
        , if model.isLoading then
            div [] [ text "Loading..." ]

          else
            case model.quote of
                Just quote ->
                    div []
                        [ text quote.content
                        , div [] [ text ("-" ++ quote.authorName) ]
                        ]

                Nothing ->
                    div [] [ text "Click the button for some wisdom!" ]
        , case model.error of
            Just err ->
                div [] [ text ("error: " ++ err) ]

            Nothing ->
                div [] [ text "" ]
        ]


main : Program String Model Msg
main =
    Browser.element
        { init = \flags -> ( init flags, Cmd.none )
        , update = update
        , subscriptions = \_ -> Sub.none
        , view = view
        }
