// scrollback: 9999999 -> set infinite scrollback
var terminal = new Terminal({
 cursorBlink: true,
 fullscreenWin: true,
 maximizeWin: true,
 screenReaderMode: true,
 scrollback: 9999999,
});
terminal.open(document.getElementById('terminal'));


// Init and load addon
var fitAddon = new FitAddon.FitAddon();
terminal.loadAddon(fitAddon);

function prompt(terminal) {
command = '';
terminal.write('\r\n$ ');
}

var frontEndCommand = {
 "clear": {
   f: () => {
     terminal.write('\x1bc')
     prompt(terminal)
   }
 }
}
var command = "";



prompt(terminal)
fitAddon.fit();

var url = "ws://webshell-ws:80/xterm"
var ws = new WebSocket(url);

ws.onclose = function(event) {
 console.log(event);
 terminal.write('\r\nconnection has been terminated from the server-side (hit refresh to restart)\r\n')
 prompt(terminal)
};

ws.onopen = function() {
 terminal._initialized = true;
 terminal.focus();
 setTimeout(function() {fitAddon.fit()});
};

ws.onmessage = (event) => {
 {{/* Process and display response from the server */}}
 if (event.data !== "") {
   terminal.write("\r\n" + event.data)
   prompt(terminal)
 }
}

terminal.onData(e => {
switch (e) {
  case '\u0003': // Ctrl+C
    terminal.write('^C');
    prompt(terminal);
    break;
  case '\r': // Enter
    if (command === '') {
     prompt(terminal);
    } else if (command in frontEndCommand) {
     frontEndCommand[command].f()
    } else {
     // Send command to server and reset command
     ws.send(command);
     command = "";
    }
    break;
  case '\u007F': // Backspace (DEL)
    // Do not delete the prompt
    if (terminal._core.buffer.x > 2) {
      terminal.write('\b \b');
      if (command.length > 0) {
        command = command.substr(0, command.length - 1);
      }
    }
    break;
  default: // Print all other characters
    if (e >= String.fromCharCode(0x20) && e <= String.fromCharCode(0x7E) || e >= '\u00a0') {
      command += e;
      terminal.write(e);
    }
}

console.log(command)
});