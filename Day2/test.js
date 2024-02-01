var s = "";
for (i = 5; i > 0; i--) {
  for (j = 0; j <= 5; j++) {
    if (j >= i) {
      s += "*";
    } else {
      s += " ";
    }
  }
  s += "\n";
}
console.log(s);