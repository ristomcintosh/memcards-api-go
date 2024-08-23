# Memcards API

## API Endpoints

Follows the API specification for [Memcards (a simple flashcard app)](https://github.com/ristomcintosh/memcards-v2). The endpoints are as follows:

### Decks

- `GET /decks` - Retrieve a list of all decks
- `GET /decks/{deckId}` - Retrieve a specific deck
- `POST /decks` - Create a new deck
- `PUT /decks/{deckId}` - Update a deck
- `DELETE /decks/{deckId}` - Delete a deck

### Flashcards

- `POST /decks/{deckId}/flashcards` - Add a new flashcard to deck
- `PUT /decks/{deckId}/flashcards/{flashcardId}` - Update a flashcard
- `DELETE /decks/{deckId}/flashcards/{flashcardId}` - Delete a flashcard
