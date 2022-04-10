package plugin

var TrelloDefaultConfig = `tools:
- name: my-trello-board
  # name of the plugin
  plugin: trello
  # options for the plugin, checkout the version from the GitHub releases
  options:
    # the repo's owner
    owner: YOUR_GITHUB_USERNAME
    # the repo's org. If you set this property, then the new repo will be created under the org you're given, and the "owner" setting above will be ignored.
    org: YOUR_ORGANIZATION_NAME
    # for which repo this board will be used
    repo: YOUR_REPO_NAME
    # the Tello board name. If empty, use owner/repo as the board's name.
    kanbanBoardName: KANBAN_BOARD_NAME`
