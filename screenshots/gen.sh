# Uses this https://github.com/homeport/termshot
# Compile and pass the executable path as first argument to this script.

TERMSHOT="$1"
CPDIFF="$2"
DIR="./screenshots"

gen() {
  cpdiff_flags=$(cat "$DIR/$1/flags" 2>/dev/null)
  output="$DIR/$1.png"
  $TERMSHOT -f $output -- "$CPDIFF $cpdiff_flags $DIR/$1/in $DIR/$1/ans" > /dev/null
  echo $2
  echo
  echo "![$1]($output)"
  echo
}

# TODO: Create one screenshot using the --show-cmd for usage (how to use it)
# it should have a pretty command such as ./my-program < in | cpdiff ans
# worst case scenario simply put it in markdown. That'd be good enough actually.
gen numbers "Comparing each number individually. Numbers are allowed to have an error."
gen heart-strings "Comparing strings. Each character is compared individually."
gen binary-strings "If you want to compare binary strings (or digits), you can compare them character by character instead of comparing their numeric value."
gen big-numbers "Big numbers are supported. Numbers can have arbitrary precision or amount of digits."
