# Learning Go

In my free time I'm having fun with Go, this is the result.

### Tips

- **Compiling and running:**

    Add this function to your ~/.bashrc

    (for amd64)

    ```bash
    # Compiling and running .go files
    go6 () {
      6g $1;
      file=${1%.go}
      6l -o $file $file'.6';
      ./$file $2;
    }
    ```

    (for 386)

    ```bash
    # Compiling and running .go files
    go8 () {
      8g $1;
      file=${1%.go}
      8l -o $file $file'.8';
      ./$file $2;
    }
    ```

    Then you can compile and run *.go* files this way:

    ```bash
    $ go6 Euler_0001.go 1000
    233168
    ```

    Then you can simply:

    ```bash
    $ ./Euler_0001 1000
    233168
    ```

