# GoSesame

A simple password manager written in Go, inspired by the magic phrase **"Open Sesame"**.  

> [!WARNING]
> This application is still in heavy development and heavily for personal use, expect things to break or things to not be perfect :)
> Consequently, this might not be the **MOST SECURE** implementation right away.

## Motivation

I am currently learning Golang and I needed some kind of password manager because creating new _secure_ passwords **and** remembering them is becoming increasingly hard.

So why not create my own? It'll be a fun challenge, a cool learning experience, and will save me a couple bucks per month! That's a win-win-win in my book.

## General Idea

I want to start simple.

This is how I envision this app to work:

1. The user provides a master password.
2. The user provides the location/use case of their password.
    - Usually the URL of a website
3. Based on user input, the app will generate a random password
    - The user can decide on:
        - the length of the password
        - the presence of uppercase characters
        - the presence of numbers
        - the presence of symbols
4. The generated password will be stored somewhere using the master password as a key to encrypt it
    - I am not sure on the encryption algorithm just yet, but it might be something like [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
    - Will most likely write to file locally or SQLite for now, might look into building some kind of SQL database to potentially allow for remote usage.
5. Once the user wants to fetch a password again, they simply need to provide their master password.

## Roadmap

- [ ] Implement base functions
  - Use master password to encrypt and decrypt passwords
  - Generate random password based on user-defined criterias
- [ ] Add a TUI
  - [bubbletea?](https://github.com/charmbracelet/bubbletea)
- [ ] Some way to use this on multiple devices at once, e.g.mobile + laptop
