# Slot Bot for Discord

## Configuration

### Environment Variables

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
