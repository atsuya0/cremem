function cremem() {
  [[ $# -eq 0 ]] \
    && print -rz $(command cremem show) \
    || command cremem $@
}

function _cremem() {
  function sub_commands() {
    _values 'Commands' \
      'remove' \
      'show'
  }

  _arguments -C \
    '(-h --help)'{-h,--help}'[help]' \
    '1: :sub_commands'
}
compdef _cremem cremem

function _cremem_hook() {
  [[ -z ${commands[cremem]} ]] && return

  local -r cmd=${1%%$'\n'}
  [[ -z ${cmd} ]] && return
  cremem register ${cmd}
}
add-zsh-hook zshaddhistory _cremem_hook
