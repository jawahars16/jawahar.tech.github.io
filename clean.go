package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/adrg/frontmatter"
)

func main() {
	updatePermaLink()
}

func updatePermaLink() {
	path := "./_posts"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		println(file.Name())
		if file.Name() == ".DS_Store" {
			continue
		}
		// bits := strings.Split(file.Name(), "-")
		// // join the bits
		// fmt.Printf("%v\n", bits)
		// title := strings.Join(bits[3:], "-")
		// title = strings.Replace(title, ".md", "", 1)

		// println(title)
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()))
		if err != nil {
			log.Default().Print(err)
			continue
		}

		content := string(data)
		content = strings.Replace(content, "permalink: ", "permalink: blog/", 1)
		// println(content)
		data = []byte(content)
		ioutil.WriteFile(fmt.Sprintf("%s/%s", path, file.Name()), data, 0644)
	}
}

func updatePosts() {

	// iterate a directory
	path := "./_posts"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		println(file.Name())
		// skip if it is file
		if !file.IsDir() {
			continue
		}

		// read file
		var matter struct {
			Date string `yaml:"date"`
		}
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/index.md", path, file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		content := string(data)
		_, err = frontmatter.Parse(strings.NewReader(content), &matter)
		// split by 'T' and get first part
		published_date := strings.Split(matter.Date, "T")[0]

		// update content
		content = strings.Replace(content, "---", "---\nlayout: post", 1)
		// update asset path using regex
		// regex to match the image path "./image.png"
		//regex := regexp.MustCompile(`./(.*).(png|jpg|jpeg)$`)
		// regex to match the image path "./image.png" and any extension - jpg, jpeg, png
		regex := regexp.MustCompile(`.\/(.*).(png|jpg|jpeg)`)
		content = regex.ReplaceAllString(content, fmt.Sprintf("/assets/blog/%s-%s/$1.$2", published_date, file.Name()))

		// convert to byte
		data = []byte(content)

		// copy file to new location
		ioutil.WriteFile(fmt.Sprintf("%s/%s-%s.md", path, published_date, file.Name()), data, 0644)
		copyImagesToAssets(fmt.Sprintf("%s/%s", path, file.Name()), fmt.Sprintf("%s-%s", published_date, file.Name()))
		// if index > 11 {
		// 	break
		// }
	}
}

func copyImagesToAssets(path string, post string) {
	// read all assets from the path
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copying assets... from %s\n", path)
	for _, asset := range files {
		if asset.Name() == "index.md" {
			continue
		}
		println(asset.Name())
		// copy to assets folder
		data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, asset.Name()))
		if err != nil {
			log.Fatal(err)
		}
		os.MkdirAll(fmt.Sprintf("./assets/blog/%s", post), os.ModePerm)
		ioutil.WriteFile(fmt.Sprintf("./assets/blog/%s/%s", post, asset.Name()), []byte(data), 0644)
	}
}
