
# Tasker

Tasker is a command-line tool written in Go for managing your todo tasks. With Tasker, you can add, update, delete, and list tasks, as well as manage their statuses (e.g., done, todo, in progress).

Project idea from <https://roadmap.sh/projects/task-tracker>

## Features

- Add new tasks
- Update existing tasks
- Delete tasks
- Clear all tasks
- List tasks (with optional filtering by status)
- Mark tasks as "done", "todo", or "in progress"
- Show task details

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/PavelNikoltsev/tasker.git
   ```

2. Navigate to the project directory:

   ```bash
   cd tasker
   ```

3. Build the project:

   ```bash
   make build
   ```

   This will generate the `tasker` executable in the project directory.

## Usage

After building the project, you can use the `tasker` command as follows:

```bash
./tasker <command> [arguments]
```

### Available Commands

| Command                               | Description                               |
| ------------------------------------- | ----------------------------------------- |
| `add <task description>`              | Add a new task                            |
| `update <task ID> <task description>` | Update an existing task                   |
| `delete <task ID>`                    | Delete a task                             |
| `clear`                               | Clear all tasks                           |
| `list [status]`                       | List tasks, optionally filtered by status |
| `show <task ID>`                      | Show details of a task                    |
| `mark-done <task ID>`                 | Mark a task as "done"                     |
| `mark-todo <task ID>`                 | Mark a task as "todo"                     |
| `mark-in-progress <task ID>`          | Mark a task as "in progress"              |
| `help`                                | Show help message                         |

### Task Statuses

You can filter tasks by their status using the `list` command:

- `done`
- `todo`
- `in-progress`

Example:

```bash
./tasker list todo
```

### Example Commands

- Add a task:
  
  ```bash
  ./tasker add "Complete the report"
  ```

- Update a task:

  ```bash
  ./tasker update 1 "Submit the report"
  ```

- Mark a task as done:

  ```bash
  ./tasker mark-done 1
  ```

- Delete a task:

  ```bash
  ./tasker delete 2
  ```

## Testing

To run the tests, use the following command:

```bash
make test
```

This will run tests in the `commands` and `models` directories.
