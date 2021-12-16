# Backend stack

## How to use

```sh
git clone https://github.com/jchonde/golang-assessment.git
cd golang-assessment
```

run it:

```sh
go run .
```

test it:

```sh
go test .
```

## What is being requested

Two endpoints need to be implemented as part of the assignment:

GET /find - Given a single word, the endpoint should find all anagrams in the english dictionary. Note that a dependency called 'word-list' is being included as part of the setup and should be used as input for all supported words in the english language.

GET /compare - Given two words, the endpoint will determine if they are anagrams or not.

##  Considerations
You may restructure the code as you see fit (while respecting the provided rules), creating new files and directories as needed. The solution to the assignment can be sent back as a zip/tar.gz file or a link to a github repo.

Things to take into account while implementing your solution:

Functionality: endpoints have to work as expected.

Design & code style: code structure, solution design, extensibility, easy to read & maintain, error handling.

Performance: code must have good performance (O(n), memory usage, etc). Multiple simultaneous calls per second should be expected for the API.

Test coverage: some level of unit testing is expected.

##  Rules
The existing setup and endpoints should be used. Endpoints should respect existing naming and query param names.

The provided word-list (english dictionary) must be used for the english language and cannot be swapped for a different one.

When getting anagrams for a given word, the passed word should not be included as part of the result. Ie: for 'cat' it should return ['act'] and not ['cat', 'act'].

The API must be served @localhost:3001 when running ```sh go run .```

## Bonus points
You can tackle as many bonus points as want/can, they are not mandatory but welcomed.

Create a new endpoint called find-longest. The endpoint should return an array containing the longest anagrams for the given dictionary. Single dictionary words (no anagrams) should be excluded.

Add support for multiple dictionaries (english remains default if not specified).

Add support to list, add and remove words from the supported dictionaries.

Add support for phrases, not just single words. Ie: 'rail safety' -> ['fairy tales'].

