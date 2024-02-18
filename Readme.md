# FASTA peptide filter

This CLI tool can be used to read input from one or multiple `FASTA` files
and extract all **peptide sequences** in a given **mass range** into a
seperate file.
Optionally, the mass of each peptide can be stored as a comment between
sequence identifier and peptide sequence.

## Notes

* all data must conform to the FASTA specification
    * this tool expects a sequence identifier to start with `>>`
    * the sequence identifier is followed by the peptide sequence
* all files are processed concurrently
* always choose a smaller/higher range limit in order to avoid float comparison
problems
    * This will be fixed in a futre (after v0.8.1) release

## Usage

Download the executable for your supported platform
[here](https://github.com/KleinSpeedy/go-peptide-filter/releases) and extract
it to a known location.

On windows it's saved as `.zip` archive, on Linux it's saved as a `tar.gz` archive.

In order to use the tool, you have to open a terminal and move to the directory
containing the CLI binary (Linux) or `.exe` (Windows).

### Examples

> [!NOTE]
> All examples are shown using Linux syntax, if you are on Windows, simply drop
> the `./` -> `peptide-analyse.exe ...` instead of `./peptide-analyse ...`.

This reads all sequences from `test.fasta` and prints all peptides with a mass between
900 to 1200 to stdout.
```sh
./peptide-analyse -s 900 -e 1200 -f test.fasta
>>test_seq_id
ELVFVPASA
```

You can also write the output into a separate file using the `>` operator.
```sh
./peptide-analyse -s 900 -e 1200 -f test.fasta > sorted.fasta
```

As of **v0.8.1** you can also write to an output file directly using the `-o` flag.
```sh
./peptide-analyse -s 900 -e 1200 -o out.fasta -f test.fasta
```

Using the `--wc` flag you can add the peptide mass as a comment between sequence
identifier and peptide sequence.
```sh
./peptide-analyse --wc -s 900 -e 1200 -f test.fasta
>>test_seq_id
; 930.84
ELVFVPASA
```

## Development

If you want to build the cli application locally:
```sh
go build .
```

Run tests locally in project directory:
```sh
go test -v ./...
```

## License

See [LICENSE](LICENSE).

## TODOs

* Mass range is checked using `>=` and `<=` operator
    * introduce better mass range check
    * check performance impact on huge datasets
    * See [here](https://floating-point-gui.de/errors/comparison/)
