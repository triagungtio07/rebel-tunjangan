package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getTotalAnak(r *bufio.Reader) (int, error) {
	input, _ := getInput("Total Anak :", reader)
	inputInt, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("harap masukkan total anak dengan format angka")
	}

	return inputInt, nil
}

func getUmurAnak(totalAnak int, r *bufio.Reader) ([]int, error) {
	var res []int
	for i := 0; i < totalAnak; i++ {
		input := fmt.Sprintf("Umur Anak Ke - %d: ", i+1)
		input2, _ := getInput(input, reader)
		umurAnakInt, err := strconv.Atoi(input2)
		if err != nil {
			return nil, errors.New("harap masukkan umur dengan format angka")
		}
		res = append(res, umurAnakInt)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(res)))

	return res, nil
}

func getGajiPokok(r *bufio.Reader) (float64, error) {
	input, _ := getInput("Gaji Pokok : Rp ", reader)
	res, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, errors.New("harap masukkan gaji dengan format angka tanpa koma dan titik")
	}
	return res, nil
}

func getTunjanganAnak(totalAnak int, umurAnak []int, gajiPokok float64) (float64, error) {
	var res float64

	if totalAnak == 1 && umurAnak[0] <= 15 {
		res += gajiPokok * 10 / 100
	} else if totalAnak > 1 {
		for i := 0; i < 2; i++ {
			if umurAnak[i] <= 5 {
				res += gajiPokok * 5 / 100
			}
			if umurAnak[i] >= 6 && umurAnak[i] <= 10 {
				res += gajiPokok * 7 / 100
			}
			if umurAnak[i] >= 11 && umurAnak[i] <= 15 {
				res += gajiPokok * 10 / 100
			}
		}

	}

	return res, nil
}

func main() {
	totalAnak, err := getTotalAnak(reader)
	if err != nil {
		println(err.Error())
		return
	}
	umurAnak, err := getUmurAnak(totalAnak, reader)
	if err != nil {
		println(err.Error())
		return
	}
	gajiPokok, err := getGajiPokok(reader)
	if err != nil {
		println(err.Error())
		return
	}
	tunjanganAnak, _ := getTunjanganAnak(totalAnak, umurAnak, gajiPokok)
	println("Tunjangan: Rp", fmt.Sprintf("%.2f", tunjanganAnak))

}
