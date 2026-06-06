<div align="center">

### 𝚜𝚑𝚊𝚛𝚎𝚗𝚟

_Sharenv is a minimal CLI tool for managing and sharing `.env` files across users and machines using AES-256-GCM encryption._

⠀· · ─────── · 𖥸 · ─────── · ·

</div>

Sharenv solves the problem of safely sharing environment variables across a team. Instead of sending secrets over Slack or storing them in plaintext, sharenv encrypts your `.env` files and stores them in a `.sharenv` lock file that is safe to commit to version control. Team members can decrypt the environments using a shared key.

**How it works:**

1. Run `sharenv init` in your project directory — creates a `.sharenv` lock file and a `.sharenv.key` encryption key.
2. Save named snapshots of your `.env` file with `sharenv save <name>`.
3. Share the key with teammates via `sharenv share`, or send them the `.sharenv.key` file directly.
4. Teammates place the key file in the project root and restore any environment with `sharenv use <name>`.

The `.sharenv` lock file holds encrypted ciphertext and is safe to commit. The `.sharenv.key` file holds the secret and **must never be committed** — add it to `.gitignore`.

## Installation

```sh
go install github.com/februu/sharenv@latest
```

Or clone and build from source:

```sh
git clone https://github.com/februu/sharenv.git
cd sharenv
go build -o sharenv .
```

## Quick Start

```sh
# 1. Initialize sharenv in your project
sharenv init

# 2. Add all key files to .gitignore immediately
echo "*.key" >> .gitignore

# 3. Commit the lock file
git add .sharenv
git commit -m "chore: add sharenv lock file"

# 4. Save your current .env as a named environment
sharenv save production

# 5. Share the key with a teammate
sharenv share
```

On a teammate's machine:
```sh
# Place the key they received into the project root as .sharenv.key, then:
sharenv use production
```

## Commands

### `sharenv init`

Initializes sharenv in the current directory. Creates two files:

- `.sharenv` — the lock file that stores encrypted environments (commit this)
- `.sharenv.key` — a 256-bit random AES key (never commit this)

```sh
sharenv init
```

> Always add `*.key` to your `.gitignore` before committing — this covers both `.sharenv.key` and the `.sharenv.old.key` backup created by `rekey`.

---

### `sharenv save <name>`

Reads the current `.env` file, encrypts it with AES-256-GCM using the local key, and stores it in `.sharenv` under the given name. If a snapshot with that name already exists, it is overwritten.

```sh
sharenv save production
sharenv save staging
sharenv save local
```

Use `-f` / `--file` to encrypt a file other than `.env`:

```sh
sharenv save production --file .env.production
```

---

### `sharenv use <name>`

Decrypts the named environment from `.sharenv` and writes it to `.env` in the current directory. Requires the correct `.sharenv.key` to be present.

```sh
sharenv use production
```

Use `-f` / `--file` to write to a file other than `.env`:

```sh
sharenv use production --file .env.production
```

If decryption fails, the key you have does not match the one used to encrypt that environment.

---

### `sharenv list`

Lists all environment names currently tracked in the `.sharenv` lock file.

```sh
sharenv list
```

---

### `sharenv remove <name>`

Removes the named environment from the `.sharenv` lock file. Does not affect the local `.env` file.

```sh
sharenv remove old-staging
```

---

### `sharenv rename <old> <new>`

Renames a tracked environment. The encrypted data is preserved — only the name changes.

```sh
sharenv rename staging qa
```

---

### `sharenv share`

Prints the contents of `.sharenv.key` to the terminal so you can copy and send it to a teammate. Handle the output carefully.

```sh
sharenv share
```

> Anyone who has this key can decrypt every environment stored in `.sharenv`.

---

### `sharenv rekey`

Generates a new 256-bit key, decrypts all tracked environments under the old key, and re-encrypts them under the new key. The old key is backed up as `.sharenv.old.key`.

```sh
sharenv rekey
```

Use this when rotating secrets or after a key has been compromised. After rekeying, distribute the new `.sharenv.key` to everyone who needs access.

## Files

| File | Commit? | Description |
|---|---|---|
| `.sharenv` | Yes | TOML lock file with base64-encoded encrypted environments |
| `.sharenv.key` | **No** | 256-bit AES secret key — treat like a password |
| `.sharenv.old.key` | **No** | Backup of the previous key, created by `rekey` |

## Security

- Encryption: AES-256-GCM (authenticated encryption — provides both confidentiality and integrity)
- Key size: 256 bits, randomly generated via `crypto/rand`
- Each encryption operation uses a unique random nonce, so encrypting the same `.env` twice produces different ciphertext
- The `.sharenv` lock file is safe to store in a public or private repository; without the key, its contents are unreadable
- The key file is written with mode `0600` (owner read/write only)
