# Chicol

Colours piped text

![Screenshot](https://github.com/chibiace/Chicol/blob/master/screenshot.png)

## Usage:

To get started, use the following command to see available options:

```bash
# prints the help
chicol -h
```

### Examples

Here we are using fortune-mod as the example input and piping it into chicol


#### Basic colours
```bash
# valid colours: black, red, green, yellow, blue, magenta, cyan, white
fortune | chicol -c red
```


#### Hex colour codes
```bash
# valid colours: any 6 digit hex code, eg. '#FFAA00'
fortune | chicol -x "#FFAA00"
```

#### Random colours
```bash
# enables random colours (default basic mode)
# basic mode uses basic colours
fortune | chicol -r
```


```bash
# ranged random mode, many more colours 
fortune | chicol -r -t range

# can also adjust the ranges with min and max flags, 
# perhaps you would like something pastel :)
fortune | chicol -r -t range -min 100 -max 255
```



#### Rainbows 
```bash
# rainbow mode
fortune | chicol -rainbow

# increase the step to cycle through the colours faster
fortune | chicol -rainbow -s 3 

# can also use ranges with rainbow
fortune | chicol -rainbow -s 8 -min 0 -max 255
```

