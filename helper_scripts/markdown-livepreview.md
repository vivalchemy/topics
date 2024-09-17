# Markdown Live Preview

## Prerequisites

- Install [pandoc](https://pandoc.org/installing.html)
- Install [zathura](https://pwmt.org/projects/zathura/)
- Install [watchexec](https://github.com/watchexec/watchexec)

## Usage

1. You will need 3 terminals open
1. In the first terminal, run `watchexec {filename} -o /tmp/pdf`
1. In the second terminal, run `zathura /tmp/pdf`
1. In the third terminal, you can edit the pdf file and it will be updated in the zathura window
