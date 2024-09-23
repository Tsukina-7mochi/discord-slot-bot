# Slot Bot for Discord

## Usage

Use docker compose to launch your bot.

```sh
$ docker compose up prod
```

## Configuration

### Environment Variables

`.env` files are supported in current working directory.

| Name              | Discription                         |
| ----------------- | ----------------------------------- |
| `TOKEN`           | Discord bot token                   |
| `APP_ID`          | Application ID of Bot               |
| `GUILD_IDS`       | Guild IDs to update commands        |
| `APP_CONFIG_FILE` | Application configuration file name |

### Application configuration

See example in [`slots.json`](./slots.json`)

- `slots`: Array of slots
  - `name`: `string`, Name of slot, used as choices of command
  - `reels`: `string[][]`, Candidate of strings for each reel
