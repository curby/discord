{{/**************************************************************************\

    YAGPDB Custom Commands - Database Management
    --------------------------------------------

    Utilities for manipulating the database. The following setup applies to
    all commands in this family.

    Read-only Commands
      dbDump This is the broadest command, and blindly dumps 100 entries at a
             time. give an option offset (100, 200, etc.) to see more "pages"
      dbSearch This searches for all entries relating to a search term,
               usually a userID or something else that relates multiple
               entries together. Run without arguments for syntax.
      dbGet This returns a single database entry value based on the given
            parameters. Run without arguments for syntax.

    Write Commands (use with care 😉)
      dbSet Sets a database entry. Run without arguments for syntax.
      dbDel Removes a database entry. Run without arguments for syntax.

    Trigger type: Command
    Trigger string: [use the command's filename]

    Setup: Restrict to only run in botspam channels
    Setup: Restrict to only admins (e.g. put in administrator cc group)

  \**************************************************************************/}}
