# ParseHelper

This is a helper to *parse text files*.
You can easily create a parser script for your use cases on top of this library.

## Usage

1. Define a new type which represents your parser.

2. Implement Parse(line string) func for the new type.

3. Implement PrintResult() func for the new type.

4. Create an instance of ParseHelper.

5. Initialize the instance, and register files and your own parsers into the helper in NewParseHelper() func of ParseHelper.

6. Execute Parse() of the instance of the ParseHelper to parse the target files.

7. Execute ShowResult() of the instance of the ParseHelper to show the result of parsing.

Please see the sample file *sample.go* for more details.

## Files

* ParseHelper.go
    * The main file.

* sample.go
    * A sample script which shows how to use ParseHelper.
