package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

//Special thanks to Robin for the scalability update

var version string = "1.9 Small bugfix"

const TotalPieces int = 60
const EndBoard int = 31

var Board = [31][54]int{
	{8, 12, 7, 1, 7, 18, 19, 4, 20, 11, 18, 8, 10, 18, 11, 20, 3, 13, 8, 16, 3, 12, 7, 20, 7, 9, 13, 5, 10, 11, 15, 13, 17, 14, 1, 10, 12, 16, 15, 10, 12, 13, 11, 16, 16, 4, 8, 14, 19, 8, 7, 8, 4, 1},
	{12, 19, 20, 15, 4, 4, 18, 19, 14, 3, 16, 7, 4, 11, 4, 12, 6, 7, 11, 5, 8, 9, 4, 8, 16, 12, 5, 3, 17, 20, 18, 18, 4, 18, 11, 13, 20, 12, 15, 4, 18, 1, 20, 1, 3, 17, 6, 5, 4, 2, 15, 5, 3, 5},
	{3, 5, 13, 5, 7, 8, 15, 19, 2, 8, 1, 3, 8, 17, 17, 7, 11, 11, 19, 19, 3, 7, 10, 10, 19, 5, 16, 4, 11, 1, 8, 18, 13, 19, 13, 4, 3, 20, 2, 18, 4, 8, 17, 19, 1, 8, 14, 15, 1, 7, 14, 18, 11, 16},
	{5, 19, 14, 17, 14, 8, 16, 12, 7, 8, 18, 7, 7, 16, 2, 19, 5, 3, 6, 5, 5, 5, 5, 6, 16, 9, 5, 4, 6, 5, 10, 6, 16, 3, 1, 8, 20, 18, 12, 14, 14, 4, 12, 9, 20, 6, 9, 16, 6, 11, 6, 20, 18, 11},
	{15, 6, 14, 15, 4, 13, 17, 6, 11, 3, 5, 15, 11, 20, 11, 10, 18, 8, 20, 12, 3, 7, 3, 15, 17, 17, 7, 3, 6, 18, 3, 1, 19, 14, 6, 5, 20, 5, 4, 9, 17, 5, 4, 8, 11, 1, 19, 17, 5, 2, 3, 11, 3, 7},
	{11, 2, 8, 7, 15, 6, 2, 1, 3, 19, 14, 14, 12, 12, 6, 1, 17, 16, 2, 16, 8, 1, 17, 4, 12, 19, 10, 9, 13, 17, 12, 6, 10, 18, 5, 6, 4, 14, 15, 6, 8, 16, 5, 19, 10, 3, 12, 10, 5, 19, 3, 19, 17, 20},
	{19, 9, 4, 2, 2, 9, 2, 4, 19, 19, 14, 13, 2, 12, 10, 17, 20, 17, 16, 12, 3, 15, 6, 19, 20, 3, 19, 17, 20, 9, 10, 9, 17, 18, 13, 14, 6, 4, 4, 19, 4, 18, 8, 1, 14, 3, 14, 14, 3, 4, 7, 11, 9, 11},
	{6, 18, 20, 1, 14, 19, 3, 19, 1, 12, 5, 5, 3, 9, 6, 20, 4, 3, 4, 2, 13, 11, 15, 19, 8, 9, 7, 5, 15, 16, 19, 2, 13, 2, 19, 10, 14, 14, 8, 16, 10, 20, 2, 4, 19, 19, 3, 12, 6, 12, 17, 1, 8, 10},
	{5, 3, 3, 13, 19, 1, 12, 5, 5, 9, 12, 12, 17, 11, 8, 3, 12, 18, 18, 13, 16, 6, 3, 15, 20, 13, 2, 9, 1, 19, 11, 11, 14, 13, 18, 13, 6, 7, 10, 18, 3, 12, 4, 10, 5, 11, 19, 16, 15, 16, 14, 9, 8, 18},
	{7, 18, 15, 12, 3, 14, 16, 19, 4, 8, 7, 6, 12, 20, 2, 11, 17, 17, 1, 1, 9, 12, 4, 13, 2, 8, 15, 19, 16, 17, 1, 12, 12, 10, 8, 13, 19, 16, 14, 17, 18, 12, 15, 19, 18, 10, 15, 2, 20, 13, 7, 16, 15, 4},
	{20, 16, 7, 14, 16, 8, 12, 5, 11, 9, 16, 13, 2, 10, 10, 19, 18, 9, 15, 12, 14, 1, 5, 3, 14, 5, 12, 18, 19, 6, 12, 9, 13, 17, 19, 15, 12, 20, 13, 9, 4, 1, 12, 13, 10, 12, 6, 6, 4, 10, 19, 14, 17, 11},
	{5, 13, 14, 20, 16, 17, 11, 19, 8, 4, 4, 10, 4, 17, 15, 12, 10, 8, 17, 10, 19, 2, 13, 9, 15, 2, 9, 14, 15, 13, 18, 10, 6, 18, 5, 18, 13, 6, 9, 13, 13, 5, 2, 13, 12, 14, 20, 3, 20, 8, 9, 12, 10, 10},
	{10, 1, 15, 3, 11, 16, 4, 5, 13, 1, 17, 12, 7, 11, 14, 14, 6, 2, 4, 1, 16, 2, 5, 13, 15, 20, 12, 17, 10, 12, 14, 15, 9, 1, 2, 20, 16, 9, 15, 4, 19, 4, 20, 7, 16, 12, 17, 18, 4, 11, 7, 11, 7, 5},
	{14, 6, 13, 2, 11, 5, 7, 1, 16, 18, 19, 9, 7, 3, 1, 14, 10, 3, 6, 4, 3, 15, 5, 5, 20, 15, 6, 5, 15, 20, 6, 1, 5, 14, 6, 14, 6, 15, 10, 18, 9, 12, 19, 5, 11, 2, 13, 4, 19, 13, 10, 5, 2, 2},
	{4, 17, 3, 9, 17, 15, 15, 8, 7, 14, 11, 17, 18, 18, 10, 7, 3, 1, 5, 5, 13, 3, 10, 17, 19, 7, 7, 20, 5, 3, 18, 7, 10, 2, 8, 17, 13, 20, 14, 7, 8, 8, 15, 5, 11, 14, 4, 7, 10, 20, 7, 20, 17, 6},
	{11, 15, 1, 18, 5, 14, 4, 2, 2, 15, 10, 11, 7, 16, 1, 4, 10, 1, 15, 3, 9, 8, 3, 10, 9, 5, 4, 1, 15, 1, 11, 11, 15, 1, 11, 12, 11, 20, 11, 12, 8, 5, 7, 19, 10, 9, 5, 5, 5, 4, 20, 2, 8, 20},
	{1, 3, 6, 10, 14, 1, 5, 4, 17, 20, 11, 5, 9, 19, 4, 5, 20, 3, 14, 12, 1, 9, 13, 11, 6, 5, 10, 19, 1, 4, 16, 13, 9, 14, 7, 13, 6, 6, 12, 16, 10, 17, 6, 13, 3, 13, 17, 6, 6, 8, 20, 20, 12, 10},
	{3, 18, 6, 13, 10, 14, 17, 8, 12, 6, 17, 17, 12, 2, 7, 4, 20, 7, 19, 17, 1, 10, 20, 12, 7, 9, 8, 1, 20, 13, 8, 15, 19, 5, 10, 14, 2, 17, 17, 7, 16, 5, 18, 11, 11, 8, 14, 8, 8, 10, 11, 20, 4, 7},
	{7, 8, 4, 11, 18, 17, 14, 12, 11, 16, 3, 7, 4, 12, 17, 12, 5, 8, 11, 19, 17, 10, 3, 14, 1, 4, 8, 17, 2, 3, 13, 18, 13, 15, 10, 7, 11, 19, 10, 1, 9, 14, 1, 11, 19, 10, 10, 4, 4, 17, 1, 5, 9, 2},
	{3, 7, 10, 8, 18, 6, 10, 9, 15, 8, 19, 16, 6, 9, 16, 11, 7, 8, 4, 4, 11, 7, 3, 14, 20, 2, 2, 13, 17, 5, 17, 7, 15, 11, 8, 13, 3, 13, 13, 6, 13, 19, 14, 10, 4, 8, 7, 11, 11, 18, 7, 10, 9, 5},
	{14, 17, 16, 8, 8, 15, 8, 8, 14, 4, 16, 9, 1, 3, 8, 1, 17, 8, 1, 10, 7, 2, 12, 19, 7, 12, 13, 14, 19, 2, 16, 8, 6, 16, 9, 17, 11, 11, 8, 20, 1, 1, 5, 7, 4, 18, 3, 15, 16, 15, 3, 11, 13, 9},
	{20, 17, 13, 16, 1, 2, 20, 13, 3, 5, 12, 2, 4, 7, 1, 12, 4, 1, 8, 19, 20, 12, 9, 6, 7, 18, 14, 1, 14, 20, 5, 6, 19, 12, 15, 11, 5, 12, 15, 9, 6, 3, 10, 1, 14, 7, 6, 14, 13, 4, 12, 9, 11, 17},
	{18, 15, 4, 20, 14, 2, 5, 13, 3, 16, 4, 4, 12, 9, 9, 4, 17, 1, 14, 20, 8, 2, 4, 10, 5, 11, 1, 18, 10, 17, 8, 2, 5, 7, 18, 1, 1, 19, 2, 16, 16, 10, 8, 18, 9, 16, 7, 7, 3, 2, 16, 6, 12, 10},
	{7, 20, 6, 13, 12, 7, 14, 10, 5, 10, 5, 11, 13, 2, 19, 9, 17, 16, 13, 16, 11, 12, 1, 9, 6, 1, 17, 12, 15, 12, 9, 13, 8, 16, 15, 7, 12, 18, 12, 11, 7, 10, 5, 19, 18, 12, 1, 19, 16, 15, 1, 1, 12, 9},
	{17, 5, 2, 6, 4, 17, 15, 7, 17, 9, 1, 19, 20, 1, 15, 13, 3, 7, 15, 17, 10, 12, 9, 7, 7, 3, 10, 19, 5, 13, 8, 18, 12, 9, 1, 12, 13, 3, 1, 9, 11, 3, 18, 7, 1, 12, 19, 14, 11, 15, 16, 2, 4, 6},
	{4, 8, 10, 2, 3, 13, 15, 4, 11, 13, 5, 9, 11, 11, 17, 18, 16, 11, 18, 18, 13, 1, 1, 5, 8, 4, 10, 17, 8, 2, 12, 5, 11, 16, 2, 16, 2, 18, 17, 20, 18, 17, 13, 14, 10, 6, 13, 15, 16, 15, 15, 2, 13, 6},
	{7, 6, 11, 3, 4, 2, 7, 8, 10, 7, 5, 3, 19, 17, 16, 9, 5, 4, 2, 17, 4, 2, 16, 4, 8, 7, 20, 8, 1, 8, 12, 16, 8, 15, 16, 20, 19, 11, 12, 14, 11, 13, 1, 2, 13, 13, 3, 18, 19, 8, 5, 14, 10, 12},
	{9, 8, 8, 16, 13, 2, 19, 5, 14, 8, 6, 2, 2, 16, 12, 11, 9, 12, 12, 1, 1, 3, 14, 14, 18, 7, 2, 2, 8, 4, 12, 17, 8, 13, 2, 1, 4, 5, 20, 2, 15, 19, 11, 7, 15, 1, 8, 6, 18, 8, 17, 17, 15, 5},
	{4, 14, 13, 20, 10, 5, 12, 12, 9, 20, 20, 13, 19, 18, 6, 14, 18, 10, 3, 16, 9, 14, 4, 2, 15, 14, 9, 13, 19, 19, 14, 20, 3, 18, 13, 17, 14, 9, 19, 14, 2, 1, 10, 12, 14, 10, 11, 8, 16, 2, 15, 9, 10, 10},
	{8, 11, 14, 7, 4, 7, 8, 17, 18, 16, 13, 8, 19, 9, 18, 15, 18, 5, 13, 2, 15, 2, 11, 14, 19, 6, 8, 19, 2, 16, 12, 6, 8, 1, 8, 7, 9, 19, 18, 16, 11, 13, 9, 15, 10, 6, 12, 11, 16, 16, 16, 10, 18, 20},
	{20, 1, 1, 1, 20, 20, 1, 1, 20, 20, 1, 1, 20, 1, 1, 1, 1, 20, 1, 20, 1, 1, 1, 20, 20, 20, 1, 20, 20, 20, 1, 1, 1, 1, 1, 1, 20, 20, 1, 1, 20, 20, 1, 20, 1, 1, 20, 1, 1, 1, 1, 20, 20, 20},
}

var PiecesSteps = [TotalPieces][2]int{
	{32, -2}, {6, -8}, {22, 2}, {4, 2}, {1, 3}, {-1, 12}, {5, 7}, {11, -2}, {-8, -2}, {20, 4}, {7, -5}, {14, 5}, {2, -3}, {8, 3}, {15, -13},
	{-4, -7}, {9, -6}, {12, -1}, {-8, -9}, {16, 5}, {17, 4}, {13, -3}, {-2, 7}, {3, 6}, {1, 10}, {-7, 3}, {5, 9}, {-4, 2}, {12, 7}, {14, 11},
	{-3, 5}, {3, -5}, {6, 5}, {8, -3}, {3, -7}, {13, 10}, {2, -10}, {1, -2}, {-2, -3}, {10, 9}, {-5, 2}, {14, 12}, {-6, -4}, {-5, 12}, {7, 2},
	{12, -6}, {18, 13}, {11, 8}, {10, -3}, {-6, 4}, {-7, -4}, {1, -7}, {9, -11}, {-10, -11}, {-15, -12}, {-1, -6}, {-11, 8}, {-9, 6}, {19, 6}, {-13, 5},
}

var Costs = [12][60]uint64{
	{40, 25, 20, 25, 10, 10, 40, 15, 5, 20, 35, 45, 20, 15, 30, 20, 15, 50, 50, 5, 35, 35, 50, 15, 30, 25, 40, 50, 15, 10, 5, 25, 45, 25, 50, 5, 10, 40, 50, 10, 40, 50, 15, 45, 30, 40, 10, 10, 45, 5, 25, 40, 15, 30, 35, 15, 20, 10, 30, 50},
	{45, 25, 10, 10, 30, 5, 50, 25, 5, 50, 25, 30, 10, 35, 50, 15, 30, 40, 10, 40, 10, 10, 30, 35, 15, 50, 10, 35, 40, 20, 35, 40, 30, 30, 30, 50, 30, 30, 40, 30, 30, 15, 45, 35, 35, 5, 25, 20, 15, 25, 15, 10, 45, 40, 25, 30, 20, 40, 25, 10},
	{15, 35, 40, 30, 45, 50, 10, 50, 15, 20, 45, 40, 20, 5, 15, 30, 50, 25, 25, 10, 15, 15, 30, 15, 5, 20, 10, 5, 40, 25, 10, 10, 45, 15, 20, 40, 30, 45, 10, 5, 10, 20, 15, 50, 15, 15, 10, 50, 50, 45, 25, 40, 50, 20, 15, 15, 10, 40, 25, 5},
	{25, 50, 35, 40, 40, 5, 50, 40, 15, 40, 25, 40, 5, 20, 25, 30, 50, 35, 45, 35, 20, 15, 10, 25, 30, 5, 10, 50, 25, 50, 50, 15, 25, 10, 10, 35, 30, 20, 25, 10, 15, 40, 20, 30, 40, 45, 35, 30, 25, 10, 10, 45, 40, 30, 30, 25, 10, 30, 45, 50},
	{40, 10, 15, 5, 5, 15, 25, 15, 15, 25, 5, 10, 45, 45, 35, 5, 25, 10, 30, 25, 40, 15, 35, 10, 40, 40, 30, 40, 50, 30, 10, 15, 40, 50, 5, 50, 50, 5, 5, 20, 45, 25, 20, 25, 30, 40, 45, 10, 10, 15, 5, 5, 35, 50, 40, 25, 15, 5, 40, 10},
	{30, 45, 15, 25, 20, 50, 15, 5, 25, 40, 50, 30, 35, 20, 35, 20, 10, 50, 35, 45, 15, 5, 20, 30, 25, 45, 15, 15, 45, 40, 40, 25, 25, 5, 20, 20, 20, 30, 35, 20, 10, 20, 40, 5, 5, 50, 15, 45, 45, 40, 45, 5, 25, 20, 25, 25, 35, 10, 15, 35},
	{10, 20, 25, 50, 40, 40, 35, 45, 5, 40, 45, 30, 10, 35, 20, 45, 35, 10, 35, 5, 50, 10, 35, 5, 35, 5, 45, 5, 40, 15, 5, 40, 10, 35, 10, 50, 25, 45, 45, 35, 40, 40, 30, 20, 30, 35, 25, 50, 25, 40, 10, 10, 20, 5, 5, 10, 50, 10, 25, 45},
	{50, 25, 10, 10, 10, 15, 15, 35, 30, 20, 50, 50, 20, 20, 45, 40, 20, 45, 50, 25, 45, 30, 30, 5, 40, 35, 45, 15, 40, 50, 45, 50, 25, 30, 30, 25, 30, 5, 5, 25, 20, 15, 25, 15, 35, 45, 50, 25, 35, 35, 10, 35, 15, 20, 30, 35, 35, 50, 20, 5},
	{35, 30, 15, 45, 5, 35, 5, 15, 30, 40, 40, 10, 40, 50, 35, 30, 50, 10, 5, 35, 10, 15, 40, 25, 40, 45, 15, 15, 15, 30, 15, 45, 10, 15, 50, 30, 15, 25, 25, 5, 10, 50, 5, 25, 20, 40, 15, 30, 40, 25, 40, 30, 50, 45, 20, 45, 40, 50, 40, 15},
	{45, 5, 20, 20, 15, 30, 45, 5, 10, 40, 45, 15, 45, 50, 25, 5, 40, 25, 30, 40, 35, 15, 50, 30, 5, 35, 35, 20, 5, 5, 5, 35, 15, 10, 25, 40, 15, 35, 20, 30, 40, 20, 45, 10, 35, 15, 35, 25, 10, 15, 10, 15, 25, 50, 40, 30, 25, 40, 15, 30},
	{40, 5, 10, 30, 15, 40, 20, 45, 40, 45, 40, 45, 50, 5, 10, 35, 45, 30, 40, 35, 5, 5, 10, 25, 5, 50, 40, 10, 35, 50, 30, 30, 10, 15, 10, 5, 35, 10, 50, 50, 30, 35, 50, 35, 50, 40, 45, 25, 5, 20, 45, 50, 10, 10, 30, 10, 10, 35, 10, 5},
	{5, 45, 15, 15, 15, 35, 35, 50, 20, 10, 45, 5, 40, 30, 50, 30, 25, 25, 20, 35, 10, 50, 15, 35, 30, 5, 20, 40, 25, 10, 30, 30, 5, 15, 15, 10, 5, 30, 30, 20, 20, 10, 10, 20, 15, 40, 45, 45, 40, 45, 15, 20, 20, 15, 10, 20, 10, 25, 40, 35},
}

var Names = [TotalPieces]string{
	"Andy", "Boris", "Caroline", "Danny", "Eric", "Fientje", "George", "Harry", "Ingrid", "Joris", "Kevin", "Leyla", "Maarten", "Nathalie", "Omar",
	"Patrick", "Quinten", "Ronny", "Sofie", "Tony", "Ursula", "Vera", "Walter", "Xavier", "Youssef", "Zoe", "Ali", "Betty", "Cody", "Dorien",
	"Eefje", "Fabiano", "Gerard", "Halima", "Ivan", "Jeanine", "Ken-Giani", "Lewis", "Mia", "Nelly", "Olly", "Polleke", "Q-Dog", "Robin", "Shelly",
	"Timmi", "Ugo", "Vladimir", "Willy", "Xing-Ming", "Yadhu", "Zena", "Gilbert", "Kristian", "Orelie", "Kobe", "Arno", "Otis", "Wannes", "Samir",
}

var Animals = [TotalPieces]string{
	"Arend", "Beer", "Cheetah", "Duif", "Egel", "Flamingo", "Giraf", "Haai", "Inktvis", "Jaguar", "Kwal", "Leeuwin", "Mol", "Neushoorn", "Olifant",
	"Panda", "Quetzal", "Rat", "Slang", "Tijger", "Uil", "Vleermuis", "Walrus", "Xenopus", "Yak", "Zeepaardje", "Aap", "Bizon", "Coyote", "Dolfijn",
	"Eekhoorn", "Fazant", "Goffer", "Haas", "Ijsbeer", "Jakhals", "Krokodil", "Luiaard", "Muis", "Nijlpaard", "Ooievaar", "Paard", "Quokka", "Raaf", "Schaap",
	"Toekan", "Usno", "Vos", "Wasbeer", "Xerus", "Yapok", "Zeehond", "Gato", "Krab", "Orca", "Koala", "Alpaca", "Otter", "Wapiti", "Stokstaartje",
}

const (
	Arend = iota
	Beer
	Cheetah
	Duif
	Egel
	Flamingo
	Giraf
	Haai
	Inktvis
	Jaguar
	Kwal
	Leeuwin
	Mol
	Neushoorn
	Olifant
	Panda
	Quetzal
	Rat
	Slang
	Tijger
	Uil
	Vleermuis
	Walrus
	Xenopus
	Yak
	Zeepaardje
	Aap
	Bizon
	Coyote
	Dolfijn
	Eekhoorn
	Fazant
	Goffer
	Haas
	Ijsbeer
	Jakhals
	Krokodil
	Luiaard
	Muis
	Nijlpaard
	Ooievaar
	Paard
	Quokka
	Raaf
	Schaap
	Toekan
	Usno
	Vos
	Wasbeer
	Xerus
	Yapok
	Zeehond
	Gato
	Krab
	Orca
	Koala
	Alpaca
	Otter
	Wapiti
	Stokstaartje
)

type CalculatedPieces struct {
	id        int
	name      string
	animal    string
	cost      uint64
	position  int
	firstPay  uint64
	totalCost uint64
}

type BasicPieces struct {
	id        int
	animal    string
	cost      uint64
	totalCost uint64
}

type Player struct {
	name        string
	human       bool
	internal    bool
	inSelection int
	inLevel     uint64
	random      bool
	bling       bool
	nodes       uint64
	cd0         int
	cd1         uint64
	moveselect  int
	totalPoints float32
}

type ScorePosition struct {
	totalCost uint64
	bm        int
}

var playerBoard, moveselect int
var bbling bool
var threads int
var playerPos int
var currNodes, gameDepth uint64
var gameTime float64

var cd0 int    //pieces
var cd1 uint64 //depth
var cd2 uint64 // nodes
var piecesCosts [TotalPieces]uint64

func initPieces(startPos int) []CalculatedPieces {
	calculatedPieces := make([]CalculatedPieces, TotalPieces)
	for i := 0; i < TotalPieces; i++ {
		calculatedPieces[i].id = i
		calculatedPieces[i].name = Names[i]
		calculatedPieces[i].animal = Animals[i]
		calculatedPieces[i].cost = piecesCosts[i]
		calculatedPieces[i].position = startPos
		calculatedPieces[i].totalCost = piecesCosts[i]
	}
	currNodes = 0
	return calculatedPieces
}

func calculateOwnPieces(ownPriorityPieces []int, board int, startPos int, bling bool) []CalculatedPieces {
	calculatedPieces := make([]CalculatedPieces, cd0)
	for i := 0; i < cd0; i++ {
		id := ownPriorityPieces[i]
		calculatedPieces[i].id = id
		calculatedPieces[i].name = Names[id]
		calculatedPieces[i].animal = Animals[id]
		calculatedPieces[i].cost = piecesCosts[id]
		calculatedPieces[i].position = startPos
		calculatedPieces[i].totalCost = piecesCosts[id]
	}
	moveselect = 0
	multiThreadingCalc(calculatedPieces, board, startPos, threads, gameDepth, bling, 0, cd0, 0, cd1)
	sortMoves(calculatedPieces)
	return calculatedPieces
}

func sortMoves(calculatedPieces []CalculatedPieces) {
	sort.Slice(calculatedPieces, func(i, j int) bool {
		return (calculatedPieces[i].totalCost) < (calculatedPieces[j].totalCost)
	})
}

func revSortMoves(calculatedPieces []CalculatedPieces) {
	sort.Slice(calculatedPieces, func(i, j int) bool {
		return (calculatedPieces[i].totalCost) > (calculatedPieces[j].totalCost)
	})
}

func calculatePieces(calculatedPieces []CalculatedPieces, board int, startPos int, gameDepth uint64, bling bool, startPieceIndex int, maxPieces int, startDepth uint64, depth uint64, nodes *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	*nodes = 0
	for p := startPieceIndex; p < maxPieces; p++ {
		var id int = calculatedPieces[p].id
		for d := startDepth; d < depth; d++ {
			if d < gameDepth/2 {
				calculatedPieces[p].position = (((calculatedPieces[p].position + PiecesSteps[id][0]) % 54) + 54) % 54
			} else {
				calculatedPieces[p].position = (((calculatedPieces[p].position + PiecesSteps[id][1]) % 54) + 54) % 54
			}
			if startDepth == 0 && d == 0 {
				calculatedPieces[p].firstPay = uint64(Board[board][calculatedPieces[p].position]) * 10
			}
			if !bling {
				if d != gameDepth/2 {
					calculatedPieces[p].totalCost += (uint64(Board[board][calculatedPieces[p].position]) * 10)
				} else {
					//changerou (double income from field)
					calculatedPieces[p].totalCost -= (uint64(Board[board][calculatedPieces[p].position]) * 20)
				}
			} else {
				if startDepth == 0 && d == 0 {
					calculatedPieces[p].totalCost = 0
				}
				if d != gameDepth/2 {
					if Board[board][calculatedPieces[p].position] > 4 {
						calculatedPieces[p].totalCost += 1
					}
					if Board[board][calculatedPieces[p].position] > 10 {
						calculatedPieces[p].totalCost += 4
					}
					if Board[board][calculatedPieces[p].position] > 17 {
						calculatedPieces[p].totalCost += 2
					}
				} else {
					//changerou (double income from field)
					if Board[board][calculatedPieces[p].position] > 4 {
						calculatedPieces[p].totalCost -= 2
					}
					if Board[board][calculatedPieces[p].position] > 10 {
						calculatedPieces[p].totalCost -= 8
					}
					if Board[board][calculatedPieces[p].position] > 17 {
						calculatedPieces[p].totalCost -= 4
					}
				}
			}
		}
		(*nodes) += (depth - startDepth)
	}
}

func multiThreadingCalc(calculatedPieces []CalculatedPieces, board int, startPos int, numThreads int, gameDepth uint64, bling bool, startPieceIndex int, maxPieces int, startDepth uint64, depth uint64) {
	var nodes = make([]uint64, numThreads)
	var threadSpread int = (int(maxPieces) - int(startPieceIndex)) / numThreads
	var newEndPiece int = startPieceIndex + threadSpread + (int(maxPieces) % numThreads)
	var wg sync.WaitGroup

	wg.Add(numThreads) //die verwacht numThreads to give wg.Done()
	for x := 0; x < numThreads; x++ {
		nodes[x] = 0
		go calculatePieces(calculatedPieces, board, startPos, gameDepth, bling, startPieceIndex, newEndPiece, startDepth, depth, &nodes[x], &wg)
		startPieceIndex = newEndPiece
		newEndPiece += threadSpread
	}

	wg.Wait() //wanneer de verwachting voltooid is ga verder
	for x := 0; x < numThreads; x++ {
		currNodes += nodes[x]
	}
}

func calculate(board int, startPos int, threads int, gameDepth uint64, bling bool, startPieceIndex int, maxPieces int, depth uint64) []CalculatedPieces {
	calculatedPieces := initPieces(startPos)
	sortMoves(calculatedPieces)
	//Degene die niet uitgerekend zijn hebben een lagere score dan degene die uitgerekend zijn
	calculatedPieces2 := make([]CalculatedPieces, maxPieces)
	for i := range calculatedPieces2 {
		calculatedPieces2[i] = calculatedPieces[i]
	}
	multiThreadingCalc(calculatedPieces2, board, startPos, threads, gameDepth, bling, 0, maxPieces, 0, depth)
	sortMoves(calculatedPieces2)
	return calculatedPieces2
}

func getAllMoves(calculatedPieces []CalculatedPieces) {
	fmt.Println("Nodes: ", currNodes)
	for i := 0; i < int(TotalPieces); i++ {
		fmt.Println(i+1, ". ", calculatedPieces[i].name, calculatedPieces[i].animal, ", Cost= ", calculatedPieces[i].cost, "First pay=", calculatedPieces[i].firstPay, ", End field=", calculatedPieces[i].position, ", Total Cost=", calculatedPieces[i].totalCost)
	}
}

func getMove(calculatedPieces []CalculatedPieces, player int, moveselection int, totalNodes uint64) {
	fmt.Println("Nodes: ", totalNodes)
	fmt.Println("Player: ", player+1)
	fmt.Println("\n", moveselection+1, ". ", calculatedPieces[moveselection].name, calculatedPieces[moveselection].animal, ", Cost= ", calculatedPieces[moveselection].cost, ", End field=", calculatedPieces[moveselection].position, ", Total Cost=", calculatedPieces[moveselection].totalCost)
}

func getPieceNR(piece string) int {
	var piecenr int
	for p := 0; p < len(Animals); p++ {
		if strings.ToLower(Animals[p]) == strings.ToLower(piece) || strings.ToLower(Names[p]) == strings.ToLower(piece) {
			piecenr = p
			break
		}
	}
	return piecenr
}

func whowins012(score1 uint64, bm1 int, score2 uint64, bm2 int) int {
	if score1 < score2 && bm2 > 1 {
		return 1
	} else if score2 < score1 && bm1 > 1 {
		return 2
	} else {
		return 0
	}
}

func whowins(piece1 int, totalCost1 uint64, bm1 int, piece2 int, totalCost2 uint64, bm2 int) {
	result := whowins012(totalCost1, bm1, totalCost2, bm2)
	if result == 1 {
		fmt.Println("\nPlayer 1 wins")
		fmt.Println(Names[piece1], Animals[piece1], "(", totalCost1, ")(bm:", bm1, ") vs",
			Names[piece2], Animals[piece2], "(", totalCost2, ")(bm:", bm2, ")")
	} else if result == 2 {
		fmt.Println("\nPlayer 2 wins")
		fmt.Println(Names[piece1], Animals[piece1], "(", totalCost1, ")(bm:", bm1, ") vs",
			Names[piece2], Animals[piece2], "(", totalCost2, ")(bm:", bm2, ")")
	} else {
		fmt.Println("\nIt's a draw")
		fmt.Println(Names[piece1], Animals[piece1], "(", totalCost1, ")(bm:", bm1, ") vs",
			Names[piece2], Animals[piece2], "(", totalCost2, ")(bm:", bm2, ")")
	}
}

func getScorePositions(board int, startPos int, gameDepth uint64, piece int) ScorePosition {
	var scorePosition ScorePosition
	var calculatedPieces = calculate(board, startPos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
	for bm := 0; bm < TotalPieces; bm++ {
		if piece == calculatedPieces[bm].id {
			scorePosition.totalCost = calculatedPieces[bm].totalCost
			scorePosition.bm = bm + 1
		}
	}
	return scorePosition
}

func setupEngineByNodes(gameDepth uint64, nodes uint64) (int, uint64, uint64) {
	var thiscd0 int
	var thiscd1 uint64
	var thiscd2 uint64
	moveselect = 0
	if gameDepth > 0 {
		if nodes > gameDepth*uint64(TotalPieces) {
			thiscd0 = TotalPieces
		} else if nodes >= gameDepth {
			temp := nodes / gameDepth
			thiscd0 = int(temp)
		} else if nodes < 10 {
			rand.Seed(time.Now().UnixNano())
			thiscd0 = TotalPieces
			pieces := int(TotalPieces - 40)
			moveselect = 40 + (rand.Intn(pieces))
		} else if nodes < gameDepth {
			thiscd0 = 1
		}
	}
	thiscd1 = gameDepth
	thiscd2 = nodes
	return thiscd0, thiscd1, thiscd2
}

func calculateWithInput() {
	//bad practice!!
	//moveselect en cd global vermijden
	var bling string
	var manual string
	fmt.Print("Manual level selection? (y/.): ")
	fmt.Scanln(&manual)
	fmt.Print("bling? (y/.): ")
	fmt.Scanln(&bling)
	if bling == "y" {
		bbling = true
	} else {
		bbling = false
	}
	if manual != "y" {
		var nodes uint64
		fmt.Print("Nodes: ")
		fmt.Scanln(&nodes)
		cd0, cd1, cd2 = setupEngineByNodes(gameDepth, nodes)
	} else {
		var pieces int
		var depth uint64
		fmt.Print("Amount pieces calculate: ")
		fmt.Scanln(&pieces)
		fmt.Print("Depth calculate: ")
		fmt.Scanln(&depth)
		fmt.Print("Move choice out amount pieces (1=best, 2=second best, ..): ")
		fmt.Scanln(&moveselect)
		moveselect--
		cd0 = pieces
		cd1 = depth
		cd2 = uint64(pieces) * depth
	}
}

func chooseInternalEngine() (int, uint64) {
	var choice int
	var level uint64
	fmt.Println("0. Level")
	fmt.Println("1. WGM Caroline")
	fmt.Println("2. Robin-BOT")
	fmt.Println("3. Oliacht")
	fmt.Println("4. Fangio")
	fmt.Println("5. Parker")
	fmt.Println("6. Confetti")
	fmt.Println("7. Frozen Frog")
	fmt.Println("8. Hector")
	fmt.Println("9. MelkMix")
	fmt.Println("10. Red Forest")
	fmt.Println("11. Fruity")
	fmt.Println("12. ExcelCheat")
	fmt.Println("13. YELLOWFISH")
	fmt.Print("Choose engine: ")
	fmt.Scanln(&choice)
	if choice == 0 {
		fmt.Print("Level: ")
		fmt.Scanln(&level)
	} else if choice == 10 {
		fmt.Print("Level (1-10): ")
		fmt.Scanln(&level)
	}
	return choice, level
}

func calculateInternalEngine(board int, startPos int, gameDepth uint64, threads int, selection int, level uint64) []CalculatedPieces {
	var bling bool
	var normalCalc bool = false
	var calculatedPieces []CalculatedPieces
	if gameDepth >= 30 {
		gameTime = 15
	} else {
		gameTime = float64(gameDepth) / 2
	}
	switch selection {
	case 0:
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		if level < 18 {
			var levelNodes = [17]uint64{0, 2, 5, 10, 20, 30, 40, 50, 60, 70, 80, 100, 120, 140, 160, 180, 200}
			cd2 = levelNodes[level-1]
		} else {
			cd2 = 200 + ((level - 17) * 40)
		}
		if cd2 < gameDepth {
			cd0 = 1
		} else {
			temp := cd2 / gameDepth
			cd0 = int(temp)
		}
		cd1 = gameDepth
		moveselect = 0
		normalCalc = true
	case 1:
		//WGM CAROLINA
		cd2 = uint64(gameTime) * 30
		cd1 = gameDepth
		temp := cd2 / gameDepth
		cd0 = int(temp)
		var priorityPieces = [53]int{
			Yak, Ijsbeer, Egel, Wasbeer, Panda, Stokstaartje, Neushoorn, Olifant, Goffer, Walrus, Coyote,
			Duif, Rat, Eekhoorn, Cheetah, Jaguar, Schaap, Flamingo, Nijlpaard, Zeehond, Krokodil,
			Toekan, Vos, Giraf, Paard, Tijger, Haas, Zeepaardje, Bizon, Slang, Raaf,
			Quokka, Xerus, Leeuwin, Muis, Aap, Beer, Dolfijn, Fazant, Haai, Mol,
			Ooievaar, Kwal, Jakhals, Usno, Luiaard, Vleermuis, Quetzal, Uil, Inktvis, Arend,
			Yapok, Xenopus,
		}
		calculatedPieces = calculateOwnPieces(priorityPieces[:], board, startPos, false)
		moveselect = 0
	case 2:
		//ROBIN
		cd2 = uint64(gameTime) * 20
		cd1 = gameDepth
		temp := cd2 / gameDepth
		cd0 = int(temp)
		var priorityPieces = [52]int{
			Krokodil, Slang, Arend, Raaf, Yapok, Ijsbeer, Zeepaardje, Goffer, Schaap, Ooievaar,
			Yak, Rat, Duif, Nijlpaard, Coyote, Panda, Vleermuis, Haai, Eekhoorn, Haas,
			Xenopus, Mol, Muis, Toekan, Vos, Quokka, Quetzal, Aap, Beer, Kwal,
			Neushoorn, Giraf, Inktvis, Xerus, Zeehond, Egel, Wasbeer, Walrus, Jakhals, Flamingo,
			Olifant, Luiaard, Paard, Cheetah, Usno, Uil, Leeuwin, Tijger, Dolfijn, Jaguar,
			Fazant, Bizon,
		}
		calculatedPieces = calculateOwnPieces(priorityPieces[:], board, startPos, false)
		moveselect = 0
	case 3:
		//OLIACHT
		cd2 = uint64(gameTime) * 15
		cd1 = gameDepth
		temp := cd2 / gameDepth
		cd0 = int(temp)
		var priorityPieces = [TotalPieces]int{
			Bizon, Cheetah, Stokstaartje, Usno, Yapok, Eekhoorn, Inktvis, Vleermuis, Uil, Yak,
			Egel, Ooievaar, Aap, Haai, Nijlpaard, Otter, Zeehond, Zeepaardje, Muis, Orca,
			Beer, Flamingo, Ijsbeer, Xenopus, Haas, Xerus, Mol, Goffer, Alpaca, Panda,
			Neushoorn, Toekan, Duif, Jakhals, Krab, Arend, Tijger, Quokka, Giraf, Schaap,
			Quetzal, Leeuwin, Luiaard, Wasbeer, Dolfijn, Raaf, Koala, Vos, Slang, Krokodil,
			Walrus, Coyote, Fazant, Gato, Rat, Olifant, Wapiti, Jaguar, Kwal, Paard,
		}
		calculatedPieces = calculateOwnPieces(priorityPieces[:], board, startPos, false)
		moveselect = 0
	case 4:
		//FANGIO
		var p float64 = 2 * gameTime
		var d float64 = 0.7 * gameTime
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		if gameDepth < 7 {
			p = (p / float64(gameDepth)) * 7
			cd0 = int(p)
			cd1 = gameDepth
		} else {
			cd0 = int(p)
			cd1 = uint64(d)
		}
		cd2 = uint64(cd0) * cd1
		moveselect = 0
		bling = true
		bbling = true
		normalCalc = true
	case 5:
		//PARKER
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd0 = 15
		cd1 = 7
		cd2 = uint64(cd0) * cd1
		moveselect = 3
		bling = false
		bbling = false
		normalCalc = true
	case 6:
		//CONFETTI
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd0 = 5
		cd1 = 6
		cd2 = uint64(cd0) * cd1
		moveselect = 1
		bling = false
		bbling = false
		normalCalc = true
	case 7:
		//FROZEN FROG
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd0 = 4
		cd1 = 2
		cd2 = uint64(cd0) * cd1
		moveselect = 0
		bling = true
		bbling = true
		normalCalc = true
	case 8:
		//HECTOR
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd0 = 20
		cd1 = 9
		cd2 = uint64(cd0) * cd1
		moveselect = 4
		bling = true
		bbling = true
		normalCalc = true
	case 9:
		//MELKMIX
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd0 = 13
		cd1 = 5
		cd2 = uint64(cd0) * cd1
		moveselect = 0
		bling = true
		bbling = true
		normalCalc = true
	case 10:
		//RED FOREST
		calculatedPieces = initPieces(startPos)
		sortMoves(calculatedPieces)
		cd2 = uint64(gameTime) * 3 * level
		temp := cd2 / gameDepth
		cd0 = int(temp)
		cd1 = gameDepth
		moveselect = 0
		bling = true
		bbling = true
		normalCalc = true
	case 11:
		//FRUITY
		var calculationPriority = [20]uint64{10, 200, 20, 190, 30, 180, 40, 170, 50, 160, 60, 150, 70, 140, 80, 130, 90, 120, 100, 110}
		if gameDepth > 30 {
			cd2 = 1000
			temp := cd2 / gameDepth
			cd0 = int(temp)
		} else {
			cd0 = 60 - int(gameDepth)
			temp := uint64(cd0) * gameDepth
			cd2 = temp
		}
		cd1 = gameDepth
		firstMoveArray := initPieces(startPos)
		calculatedPieces = make([]CalculatedPieces, cd0)
		multiThreadingCalc(firstMoveArray, board, startPos, threads, gameDepth, false, 0, TotalPieces, 0, 1)
		piececounter := 0
		for i := range calculationPriority {
			for j := 0; j < int(TotalPieces); j++ {
				if piececounter == int(cd0) {
					goto end
				} else if firstMoveArray[j].firstPay == calculationPriority[i] {
					calculatedPieces[piececounter] = firstMoveArray[j]
					piececounter++
				}
			}
		}
	end:
		moveselect = 0
		calculatedPieces2 := make([]CalculatedPieces, cd0)
		for i := range calculatedPieces2 {
			calculatedPieces2[i] = calculatedPieces[i]
		}
		multiThreadingCalc(calculatedPieces2, board, startPos, threads, gameDepth, bling, 0, cd0, 1, cd1)
		sortMoves(calculatedPieces2)
		return calculatedPieces2
	case 12:
		//EXCEL CHEAT
		cd2 = uint64(gameTime) * 20
		cd1 = gameDepth
		temp := cd2 / gameDepth
		cd0 = int(temp)
		var priorityPieces = [TotalPieces]int{
			Usno, Zeehond, Luiaard, Egel, Yak, Krokodil, Mol, Ijsbeer, Fazant, Xenopus,
			Duif, Giraf, Aap, Beer, Goffer, Kwal, Schaap, Flamingo, Koala, Walrus,
			Muis, Eekhoorn, Bizon, Panda, Raaf, Ooievaar, Xerus, Quokka, Zeepaardje, Yapok,
			Inktvis, Slang, Haas, Neushoorn, Gato, Quetzal, Wasbeer, Nijlpaard, Haai, Vos,
			Toekan, Rat, Coyote, Otter, Krab, Alpaca, Stokstaartje, Orca, Vleermuis, Jakhals,
			Leeuwin, Dolfijn, Paard, Olifant, Tijger, Uil, Wapiti, Jaguar, Cheetah, Arend,
		}
		calculatedPieces = calculateOwnPieces(priorityPieces[:], board, startPos, false)
		moveselect = 0
	case 13:
		//YELLOWFISH
		calculatedPieces = initPieces(startPos)
		revSortMoves(calculatedPieces)
		if gameDepth < 10 {
			cd0 = 50 - int(gameDepth)
			temp := uint64(cd0) * gameDepth
			cd2 = temp
		} else if gameDepth < 20 {
			cd0 = 40
			temp := uint64(cd0) * gameDepth
			cd2 = temp
		} else {
			cd2 = 800
			temp := cd2 / gameDepth
			cd0 = int(temp)
		}
		cd1 = gameDepth
		cd2 = uint64(cd0) * cd1
		moveselect = 0
		bling = false
		normalCalc = true
	}

	if normalCalc {
		calculatedPieces2 := make([]CalculatedPieces, cd0)
		for i := range calculatedPieces2 {
			calculatedPieces2[i] = calculatedPieces[i]
		}
		multiThreadingCalc(calculatedPieces2, board, startPos, threads, gameDepth, bling, 0, cd0, 0, cd1)
		sortMoves(calculatedPieces2)
		return calculatedPieces2
	}
	return calculatedPieces
}

func animalPerf(gStart int, gEnd int, gameDepth uint64, eDepth uint64, trap uint64, lowestnodes uint64, highestnodes uint64, bStart int, bEnd int, startPos int, endPos int, sTop int, eTop int, random uint64, startR uint64, stopR uint64, calculated *uint64, piecesArray *[]BasicPieces, totalNodesArray *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	var cdTemp0 int
	var cdTemp1, cdTemp2 uint64
	for ; gameDepth <= eDepth; gameDepth += trap {
		for curRat := lowestnodes; curRat <= highestnodes; curRat++ {
			cdTemp0, cdTemp1, cdTemp2 = setupEngineByNodes(gameDepth, curRat)
			if random == 0 {
				for g := gStart; g <= gEnd; g++ {
					pickFirstCost(g, gameDepth)
					for curBoard := bStart; curBoard <= bEnd; curBoard++ {
						for pos := startPos; pos <= endPos; pos++ {
							calculatedPieces := calculate(curBoard, pos, 1, gameDepth, false, 0, cdTemp0, cdTemp1)
							(*calculated)++
							top := sTop
							for ; top <= eTop; top++ {
								id := calculatedPieces[top].id
								(*piecesArray)[id].totalCost += (uint64(eTop) + 1) - uint64(top)
							}
							(*totalNodesArray) += cdTemp2
						}
					}
				}
			} else {
				for i := startR; i < stopR; i++ {
					rand.Seed(time.Now().UnixNano())
					curBoard := rand.Intn(EndBoard)
					pos := rand.Intn(54)
					generation := rand.Intn(12)
					pickFirstCost(generation+1, gameDepth)
					calculatedPieces := calculate(curBoard, pos, 1, gameDepth, false, 0, cdTemp0, cdTemp1)
					(*calculated)++
					top := sTop
					for ; top <= eTop; top++ {
						id := calculatedPieces[top].id
						(*piecesArray)[id].totalCost += (uint64(eTop) + 1) - uint64(top)
					}
					(*totalNodesArray) += cdTemp2
				}
			}
		}
	}
}

func firstScoreAnimalPerf(startGeneration int, endGeneration int, gameDepth uint64, eDepth uint64, trap uint64, lowestnodes uint64, highestnodes uint64, bStart int, bEnd int, startPos int, endPos int, sTop int, eTop int, random uint64, startR uint64, stopR uint64, calculated *uint64, piecesArray *[][]BasicPieces, totalNodesArray *uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	var cdTemp0 int
	var cdTemp1, cdTemp2 uint64
	for ; gameDepth <= eDepth; gameDepth += trap {
		for g := startGeneration; g <= endGeneration; g++ {
			for curRat := lowestnodes; curRat <= highestnodes; curRat++ {
				pickFirstCost(g, gameDepth)
				cdTemp0, cdTemp1, cdTemp2 = setupEngineByNodes(gameDepth, curRat)
				if random == 0 {
					for curBoard := bStart; curBoard <= bEnd; curBoard++ {
						for pos := startPos; pos <= endPos; pos++ {
							calculatedPieces := calculate(curBoard, pos, 1, gameDepth, false, 0, cdTemp0, cdTemp1)
							(*calculated)++
							top := sTop
							for ; top <= eTop; top++ {
								id := calculatedPieces[top].id
								(*piecesArray)[(Board[curBoard][pos])-1][id].totalCost += (uint64(eTop) + 1) - uint64(top)
							}
							(*totalNodesArray) += cdTemp2
						}
					}
				} else {
					for i := startR; i < stopR; i++ {
						rand.Seed(time.Now().UnixNano())
						curBoard := rand.Intn(EndBoard)
						pos := rand.Intn(54)
						calculatedPieces := calculate(curBoard, pos, 1, gameDepth, false, 0, cdTemp0, cdTemp1)
						(*calculated)++
						top := sTop
						for ; top <= eTop; top++ {
							id := calculatedPieces[top].id
							(*piecesArray)[(Board[curBoard][pos])-1][id].totalCost += (uint64(eTop) + 1) - uint64(top)
						}
						(*totalNodesArray) += cdTemp2
					}
				}
			}
		}
	}
}

func animalCostsMultiThreading(startGeneration int, endGeneration int, startDepth uint64, endDepth uint64, depthSteps uint64, startBoard int, endBoard int, startPos int, endPos int, animalCostsHits *[]float64, totalHits *float64, wg *sync.WaitGroup) {
	defer wg.Done()
	d := startDepth
	for ; d <= endDepth; d += depthSteps {
		for g := startGeneration; g <= endGeneration; g++ {
			pickFirstCost(g, d)
			for b := startBoard; b <= endBoard; b++ {
				for pos := startPos; pos <= endPos; pos++ {
					calculatedPieces := calculate(b, pos, 1, uint64(d), false, 0, TotalPieces, uint64(d))
					(*animalCostsHits)[(calculatedPieces[0].cost/(uint64(d)*5))-1]++
					*totalHits++
				}
			}
		}
	}
}

func calculateFirstCost() {
	rand.Seed(time.Now().UnixNano())
	for j := range piecesCosts {
		r := rand.Intn(10)
		piecesCosts[j] = (gameDepth * 5) + (uint64(r) * (gameDepth * 5))
	}
}

func pickFirstCost(generation int, gameDepth uint64) {
	for i := range piecesCosts {
		piecesCosts[i] = (Costs[generation-1][i]) * gameDepth
	}
}

func getAnimals() {
	for i := 0; i < len(piecesCosts)/2; i++ {
		fmt.Println(Names[i], Animals[i])
	}
	fmt.Println("\n\n")
	for j := len(piecesCosts) / 2; j < len(piecesCosts); j++ {
		fmt.Println(Names[j], Animals[j])
	}
}

func getAnimalsCosts() {
	fmt.Println("\n\n")
	fmt.Print("{", piecesCosts[0])
	for a := 1; a < len(piecesCosts); a++ {
		fmt.Print(", ", piecesCosts[a])
	}
	fmt.Print("},\n")
	for i := 0; i < len(piecesCosts)/2; i++ {
		fmt.Println(piecesCosts[i])
	}
	fmt.Println("\n")
	for j := len(piecesCosts) / 2; j < len(piecesCosts); j++ {
		fmt.Println(piecesCosts[j])
	}
}

func createCostsGenerations(generations int) {
	originalGameDepth := gameDepth
	gameDepth = 1
	for i := 0; i < generations; i++ {
		calculateFirstCost()
		getAnimalsCosts()
	}
	gameDepth = originalGameDepth
}

func main() {
	var choice int
	threads = 1
	for {
		fmt.Println("\n0. Show version")
		fmt.Println("1. Change board")
		fmt.Println("2. Change position")
		fmt.Println("3. Random cost generation/or board/position")
		fmt.Println("4. Change gamedepth & cost generation")
		fmt.Println("5. Amount threads")
		fmt.Println("6. Overview moves")
		fmt.Println("7. Calculate engine move")
		fmt.Println("8. Play vs Engine")
		fmt.Println("9. Engine vs Engine")
		fmt.Println("10. Engine nodes vs Engine nodes")
		fmt.Println("11. 1v1")
		fmt.Println("12. Benchmark")
		fmt.Println("13. Animals performance list")
		fmt.Println("14. Calculate performance nodes from games")
		fmt.Println("15. Multiple players/rounds game")
		fmt.Println("16. Overview difficult level positions")
		fmt.Println("17. Animal picked by nodes/engine for different positions")
		fmt.Println("18. Analyze board")
		fmt.Println("19. In what positions is animal X good")
		fmt.Println("20. Animal Cost winning percentage")
		fmt.Println("21. Start field statistics")
		fmt.Println("22. Show animals")
		fmt.Println("23. Generate costs generations")
		fmt.Print("\nChoose nr: ")
		fmt.Scanln(&choice)
		var generation int

		if choice == 0 {
			fmt.Println(version)
		}

		if choice == 1 {
			fmt.Print("\nBoardNR: ")
			fmt.Scanln(&playerBoard)
			playerBoard--
		}

		if choice == 2 {
			fmt.Print("Players startposition: ")
			fmt.Scanln(&playerPos)
		}

		if choice == 3 {
			var randomBoard string
			rand.Seed(time.Now().UnixNano())
			fmt.Print("Random Board? (y/.): ")
			fmt.Scanln(&randomBoard)
			if randomBoard == "y" {
				playerBoard = rand.Intn(EndBoard)
			} else {
				fmt.Print("Board: ")
				fmt.Scanln(&playerBoard)
				playerBoard--
			}
			playerPos = rand.Intn(54) // van 0 tot 54 dus tot en met 53
			if gameDepth == 0 {
				fmt.Print("Game Depth: ")
				fmt.Scanln(&gameDepth)
			}
			generation = rand.Intn(12) + 1
			pickFirstCost(generation, gameDepth)
			fmt.Println("\nAnimals Cost Generation: ", generation)
			fmt.Println("Board: ", playerBoard+1)
			fmt.Println("Players startposition: ", playerPos)
		}

		if choice == 4 {
			fmt.Print("Game depth (total moves to play with a piece): ")
			fmt.Scanln(&gameDepth)
			if gameDepth > 30 {
				gameTime = 15
			} else {
				gameTime = float64(gameDepth) / 2
			}
			fmt.Println("Game time ", gameTime, " minutes")
			fmt.Print("Costs Generation: ")
			fmt.Scanln(&generation)
			pickFirstCost(generation, gameDepth)
		}

		if choice == 5 {
			fmt.Print("Amount of threads: ")
			fmt.Scanln(&threads)
		}

		if choice == 6 {
			var bbling bool
			var bling string
			fmt.Print("Bling? (y/.): ")
			fmt.Scanln(&bling)
			if bling == "y" {
				bbling = true
			} else {
				bbling = false
			}
			if gameDepth == 0 {
				fmt.Print("Game Depth: ")
				fmt.Scanln(&gameDepth)
				pickFirstCost(generation, gameDepth)
			}
			fmt.Print("\n OVERVIEW -->")
			calculatedPieces := calculate(playerBoard, playerPos, threads, gameDepth, bbling, 0, TotalPieces, gameDepth)
			getAllMoves(calculatedPieces)
		}

		if choice == 7 || choice == 8 {
			var internal string
			var calculatedPieces []CalculatedPieces
			fmt.Print("Internal engine? (y/.): ")
			fmt.Scanln(&internal)
			if internal != "y" {
				calculateWithInput()
				calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, bbling, 0, cd0, cd1)
			} else {
				selection, level := chooseInternalEngine()
				calculatedPieces = calculateInternalEngine(playerBoard, playerPos, gameDepth, threads, selection, level)
			}
			currN := currNodes
			if choice == 7 {
				getMove(calculatedPieces, 0, moveselect, currN)
			} else if choice == 8 {
				var cpiece1, cpiece2 string
				var piece1, piece2 int
				cpiece1 = calculatedPieces[moveselect].animal
				fmt.Print("Your move --> animal: ")
				fmt.Scanln(&cpiece2)
				piece1 = getPieceNR(cpiece1)
				piece2 = getPieceNR(cpiece2)

				scorePosition1 := getScorePositions(playerBoard, playerPos, gameDepth, piece1)
				scorePosition2 := getScorePositions(playerBoard, playerPos, gameDepth, piece2)
				fmt.Println("\nNodes: ", currN)
				fmt.Println("Engine AmountPieces: ", cd0, " | Engine Depth: ", cd1)
				whowins(piece1, scorePosition1.totalCost, scorePosition1.bm, piece2, scorePosition2.totalCost, scorePosition2.bm)
			}
		}

		if choice == 9 {
			var piece1, piece2, amountPieces1, amountPieces2 int
			var depth1, depth2 uint64
			var bling1, bling2, manual1, manual2, internal1, internal2 string
			var bbling1, bbling2 bool
			bbling1 = false
			bbling2 = false
			var calculatedPieces []CalculatedPieces

			fmt.Print("Engine 1 --> Internal engine? (y/.): ")
			fmt.Scanln(&internal1)
			if internal1 == "y" {
				selection, level := chooseInternalEngine()
				calculatedPieces = calculateInternalEngine(playerBoard, playerPos, gameDepth, threads, selection, level)
				amountPieces1 = cd0
				depth1 = cd1
			} else {
				fmt.Print("Engine 1 --> Manual level selection? (y/.): ")
				fmt.Scanln(&manual1)
				fmt.Print("Engine 1 --> bling? (y/.): ")
				fmt.Scanln(&bling1)
				if bling1 == "y" {
					bbling1 = true
				}
				if manual1 != "y" {
					var nodes1 uint64
					fmt.Print("Nodes: ")
					fmt.Scanln(&nodes1)
					cd0, cd1, cd2 = setupEngineByNodes(gameDepth, nodes1)
					amountPieces1 = cd0
					depth1 = cd1
					moveselect = 0
				} else {
					fmt.Print("Engine 1 --> Amount pieces calculate: ")
					fmt.Scanln(&amountPieces1)
					fmt.Print("Engine 1 --> Depth calculate: ")
					fmt.Scanln(&depth1)
					fmt.Print("Engine 1 --> Move choice out amount pieces (1=best, 2=second best, ..): ")
					fmt.Scanln(&moveselect)
					moveselect--
				}
				calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, bbling1, 0, amountPieces1, depth1)
			}
			name1 := calculatedPieces[moveselect].animal
			piece1 = getPieceNR(name1)
			scorePosition1 := getScorePositions(playerBoard, playerPos, gameDepth, piece1)
			fmt.Print("Engine 2 --> Internal engine? (y/.): ")
			fmt.Scanln(&internal2)
			if internal2 == "y" {
				selection, level := chooseInternalEngine()
				calculatedPieces = calculateInternalEngine(playerBoard, playerPos, gameDepth, threads, selection, level)
				amountPieces2 = cd0
				depth2 = cd1
			} else {
				fmt.Print("Engine 2 --> Manual level selection? (y/.): ")
				fmt.Scanln(&manual2)
				fmt.Print("Engine 2 --> bling? (y/.): ")
				fmt.Scanln(&bling2)
				if bling2 == "y" {
					bbling2 = true
				}
				if manual2 != "y" {
					var nodes2 uint64
					fmt.Print("Nodes: ")
					fmt.Scanln(&nodes2)
					cd0, cd1, cd2 = setupEngineByNodes(gameDepth, nodes2)
					amountPieces2 = cd0
					depth2 = cd1
					moveselect = 0
				} else {
					fmt.Print("Engine 2 --> Amount pieces calculate: ")
					fmt.Scanln(&amountPieces2)
					fmt.Print("Engine 2 --> Depth calculate: ")
					fmt.Scanln(&depth2)
					fmt.Print("Engine 2 --> Move choice out amount pieces (1=best, 2=second best, ..): ")
					fmt.Scanln(&moveselect)
					moveselect--
				}
				calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, bbling2, 0, amountPieces2, depth2)
			}
			var name2 string = calculatedPieces[moveselect].animal
			piece2 = getPieceNR(name2)
			scorePosition2 := getScorePositions(playerBoard, playerPos, gameDepth, piece2)

			fmt.Println("\nAmountPieces player 1: ", amountPieces1, " | Depth player 1: ", depth1)
			fmt.Println("AmountPieces player 2: ", amountPieces2, " | Depth player 2: ", depth2)
			whowins(piece1, scorePosition1.totalCost, scorePosition1.bm, piece2, scorePosition2.totalCost, scorePosition2.bm)
		}

		if choice == 10 {
			var depth1, depth2, nodes1, nodes2 uint64
			var piece1, piece2, amountPieces1, amountPieces2 int
			fmt.Print("Engine 1 --> Nodes: ")
			fmt.Scanln(&nodes1)
			cd0, cd1, cd2 = setupEngineByNodes(gameDepth, nodes1)
			calculatedPieces := calculate(playerBoard, playerPos, threads, gameDepth, false, 0, cd0, cd1)
			name1 := calculatedPieces[0].animal
			piece1 = getPieceNR(name1)
			amountPieces1 = cd0
			depth1 = cd1

			fmt.Print("Engine 2 --> Nodes: ")
			fmt.Scanln(&nodes2)
			cd0, cd1, cd2 = setupEngineByNodes(gameDepth, nodes2)
			calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, false, 0, cd0, cd1)
			var name2 string = calculatedPieces[0].animal
			piece2 = getPieceNR(name2)
			amountPieces2 = cd0
			depth2 = cd1

			scorePosition1 := getScorePositions(playerBoard, playerPos, gameDepth, piece1)
			scorePosition2 := getScorePositions(playerBoard, playerPos, gameDepth, piece2)

			fmt.Println("\nAmountPieces player 1: ", amountPieces1, " | Depth player 1: ", depth1)
			fmt.Println("AmountPieces player 2: ", amountPieces2, " | Depth player 2: ", depth2)
			whowins(piece1, scorePosition1.totalCost, scorePosition1.bm, piece2, scorePosition2.totalCost, scorePosition2.bm)
		}

		if choice == 11 {
			var cpiece1, cpiece2 string
			var piece1, piece2 int
			fmt.Print("Player 1 --> animal: ")
			fmt.Scanln(&cpiece1)
			fmt.Print("Player 2 --> animal: ")
			fmt.Scanln(&cpiece2)
			piece1 = getPieceNR(cpiece1)
			piece2 = getPieceNR(cpiece2)
			scorePosition1 := getScorePositions(playerBoard, playerPos, gameDepth, piece1)
			scorePosition2 := getScorePositions(playerBoard, playerPos, gameDepth, piece2)
			whowins(piece1, scorePosition1.totalCost, scorePosition1.bm, piece2, scorePosition2.totalCost, scorePosition2.bm)
		}

		if choice == 12 {
			var pieces int
			var depth uint64
			fmt.Print("Pieces to calculate: ")
			fmt.Scanln(&pieces)
			fmt.Print("Depth to calculate: ")
			fmt.Scanln(&depth)
			if gameDepth == 0 {
				fmt.Print("Game Depth: ")
				fmt.Scanln(&gameDepth)
			}
			generation = rand.Intn(12) + 1
			pickFirstCost(generation, gameDepth)
			fmt.Println("Engine choose as player 1 --> ")
			start := time.Now()
			calculatedPieces := calculate(playerBoard, playerPos, threads, gameDepth, false, 0, pieces, depth)
			getMove(calculatedPieces, 0, 0, currNodes)
			fmt.Println("Time: ", time.Since(start))
		}

		if choice == 13 {
			var originalPos0 int = playerPos
			var originalGameDepth uint64 = gameDepth
			var lowestnodes, highestnodes, eDepth uint64
			var trap uint64 = 1
			piecesArray := make([][]BasicPieces, threads)
			for t := 0; t < threads; t++ {
				piecesArray[t] = make([]BasicPieces, TotalPieces)
			}
			var gStart, gEnd, bStart, bEnd, pStart, pEnd, sTop, eTop int
			var randoms string
			var random uint64 = 0
			for a := 0; a < threads; a++ {
				for i := 0; i < TotalPieces; i++ {
					piecesArray[a][i].id = i
					piecesArray[a][i].animal = Animals[i]
					piecesArray[a][i].totalCost = 0
				}
			}
			fmt.Print("Random position? (y/.): ")
			fmt.Scanln(&randoms)
			if randoms == "y" {
				fmt.Print("How many random positions per depth: ")
				fmt.Scanln(&random)
			} else {
				fmt.Print("Pieces Costs Generation Start: ")
				fmt.Scanln(&gStart)
				fmt.Print("Pieces Costs Generation End: ")
				fmt.Scanln(&gEnd)
				fmt.Print("Start from board: ")
				fmt.Scanln(&bStart)
				fmt.Print("To board: ")
				fmt.Scanln(&bEnd)
				fmt.Print("Start from position: ")
				fmt.Scanln(&pStart)
				fmt.Print("To position: ")
				fmt.Scanln(&pEnd)
			}
			fmt.Print("Start game depth: ")
			fmt.Scanln(&gameDepth)
			fmt.Print("End game depth: ")
			fmt.Scanln(&eDepth)
			fmt.Print("Depth in steps of (1,2,..): ")
			fmt.Scanln(&trap)
			fmt.Print("Top: ")
			fmt.Scanln(&sTop)
			fmt.Print("Bottom: ")
			fmt.Scanln(&eTop)
			fmt.Print("Lowest nodes: ")
			fmt.Scanln(&lowestnodes)
			fmt.Print("Highest nodes: ")
			fmt.Scanln(&highestnodes)
			bStart--
			bEnd--
			sTop--
			eTop--
			currNodes = 0
			calculated := make([]uint64, threads)
			totalNodesArray := make([]uint64, threads)
			var threadSpread int = (pEnd - pStart) / threads
			var newEndPos int = pStart + threadSpread + ((pEnd - pStart) % threads)
			startPos := make([]int, threads)
			endPos := make([]int, threads)
			endPos[0] = (((pStart + newEndPos) % 54) + 54) % 54
			for a := 1; a < threads; a++ {
				startPos[a] = (((endPos[a-1] + 1) % 54) + 54) % 54
				endPos[a] = (((endPos[a-1] + threadSpread) % 54) + 54) % 54
			}

			var wg sync.WaitGroup
			wg.Add(threads)
			var startR uint64 = 0
			var stopR uint64 = (random / uint64(threads)) + (random % uint64(threads))
			for t := 0; t < threads; t++ {
				if randoms != "y" {
					go animalPerf(gStart, gEnd, gameDepth, eDepth, trap, lowestnodes, highestnodes, bStart, bEnd, startPos[t], endPos[t], sTop, eTop, random, 0, 0, &calculated[t], &piecesArray[t], &totalNodesArray[t], &wg)
				} else {
					go animalPerf(1, 1, gameDepth, eDepth, trap, lowestnodes, highestnodes, bStart, bEnd, startPos[t], endPos[t], sTop, eTop, random, startR, stopR, &calculated[t], &piecesArray[t], &totalNodesArray[t], &wg)
					startR = stopR
					stopR += (random / uint64(threads))
				}
			}
			wg.Wait()

			var totalNodes uint64
			var totalCalculated uint64
			pieces := make([]BasicPieces, TotalPieces)
			for p := 0; p < TotalPieces; p++ {
				pieces[p].id = p
				pieces[p].animal = Animals[p]
				pieces[p].totalCost = 0
			}

			for x := 0; x < threads; x++ {
				totalNodes += totalNodesArray[x]
				totalCalculated += calculated[x]
				for p := 0; p < int(TotalPieces); p++ {
					id := piecesArray[x][p].id
					pieces[id].totalCost += piecesArray[x][p].totalCost
				}
			}

			sort.Slice(pieces, func(i, j int) bool {
				return (pieces[i].totalCost) > (pieces[j].totalCost)
			})

			fmt.Println("\nTotal Nodes: ", totalNodes)
			fmt.Println("Total position calculations: ", totalCalculated)
			var current int = 1
			for i := range pieces {
				if pieces[i].totalCost > 0 {
					fmt.Println(current, pieces[i].animal, pieces[i].totalCost)
					current++
				}
			}
			playerPos = originalPos0
			gameDepth = originalGameDepth
		}

		if choice == 14 {
			var originalGameDepth uint64 = gameDepth
			var originalBoard int = playerBoard
			var originalPiecesCosts = piecesCosts
			var amountgames int
			var averagePositionNodes uint64 = 0
			var averageUpperNodes uint64 = 0
			var generation int
			var board int
			var pos int
			fmt.Print("Amount of games: ")
			fmt.Scanln(&amountgames)
			for i := 0; i < amountgames; i++ {
				fmt.Print("Pieces Costs Generation: ")
				fmt.Scanln(&generation)
				fmt.Print("Board: ")
				fmt.Scanln(&board)
				playerBoard = board - 1
				fmt.Print("Game Depth: ")
				fmt.Scanln(&gameDepth)
				pickFirstCost(generation, gameDepth)
				fmt.Print("Position: ")
				fmt.Scanln(&pos)
				var curAnimal string
				var curIdAnimal int
				fmt.Print("Animal in that game: ")
				fmt.Scanln(&curAnimal)
				curIdAnimal = getPieceNR(curAnimal)
				calculatedPieces := calculate(playerBoard, pos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
				index := 0
				for i := range calculatedPieces {
					if curIdAnimal == calculatedPieces[i].id {
						break
					} else {
						index++
					}
				}
				var score uint64 = calculatedPieces[index].totalCost
				var findedOwnScore bool = false
				var j uint64 = gameDepth
				for ; j <= gameDepth*uint64(TotalPieces); j += gameDepth {
					temp := j / gameDepth
					cd0, cd1 = int(temp), gameDepth
					calculatedPieces = calculate(playerBoard, pos, threads, gameDepth, false, 0, cd0, cd1)
					if calculatedPieces[0].totalCost < score {
						fmt.Println(">", j)
						averageUpperNodes += j
						if !findedOwnScore && j-gameDepth > 0 {
							fmt.Println("=", j-(gameDepth/2))
							averagePositionNodes += (j - (gameDepth / 2))
							findedOwnScore = true
						}
						goto done
					}
					if !findedOwnScore && (calculatedPieces[0].id == curIdAnimal || calculatedPieces[0].totalCost == score) {
						fmt.Println("=", j)
						averagePositionNodes += j
						findedOwnScore = true
					}
					if j == gameDepth*uint64(TotalPieces) && calculatedPieces[0].totalCost == score {
						fmt.Println(">", j+gameDepth)
						averageUpperNodes += j + gameDepth
					}
				}
			done:
			}
			var fAverPos float64 = float64(averagePositionNodes) / float64(amountgames)
			var fAverUpp float64 = float64(averageUpperNodes) / float64(amountgames)
			fmt.Println("\nAverage nodes to find your solutions is: ", fAverPos)
			fmt.Println("Average nodes to find better solutions is: ", fAverUpp)
			fmt.Println("Median nodes is: ", (fAverPos+fAverUpp)/2)
			gameDepth = originalGameDepth
			playerBoard = originalBoard
			piecesCosts = originalPiecesCosts
			moveselect = 0
		}

		if choice == 15 {
			var amountplayers, rounds, generation int
			var handmatig, internal, randomBoard, manualDepth, manualGeneration string
			fmt.Print("How many players on 1 board: ")
			fmt.Scanln(&amountplayers)
			fmt.Print("How many rounds: ")
			fmt.Scanln(&rounds)
			fmt.Print("Manual Pieces Costs Generation each round? (y/.): ")
			fmt.Scanln(&manualGeneration)
			fmt.Print("Manual gamedepth each round? (y/.): ")
			fmt.Scanln(&manualDepth)
			if manualDepth != "y" {
				fmt.Print("Gamedepth: ")
				fmt.Scanln(&gameDepth)
			}
			fmt.Print("Random board? (y/.): ")
			fmt.Scanln(&randomBoard)
			if randomBoard != "y" {
				fmt.Print("Board: ")
				fmt.Scanln(&playerBoard)
				playerBoard--
				fmt.Print("Manual position input each round? (y/.): ")
				fmt.Scanln(&handmatig)
			}
			var pieces = make([]CalculatedPieces, amountplayers)
			var bmNumbers = make([]int, amountplayers)
			var players = make([]Player, amountplayers)
			for player := 0; player < amountplayers; player++ {
				var name string
				var humanS string
				scanner := bufio.NewScanner(os.Stdin)
				fmt.Print("Name of current player: ")
				scanner.Scan()
				name = scanner.Text()
				players[player].name = name
				fmt.Print("Internal engine? (y/.): ")
				fmt.Scanln(&internal)
				if internal == "y" {
					players[player].internal = true
					selection, level := chooseInternalEngine()
					players[player].inSelection = selection
					players[player].inLevel = level
				} else {
					fmt.Print("Is current player human? (y/.): ")
					fmt.Scanln(&humanS)
					if humanS == "y" {
						players[player].human = true
					} else {
						calculateWithInput()
						players[player].cd0 = cd0
						players[player].cd1 = cd1
						players[player].nodes = cd2
						players[player].moveselect = moveselect
						players[player].bling = bbling
						if moveselect > 39 {
							players[player].random = true
						}
					}
				}
			}

			rand.Seed(time.Now().UnixNano())
			var calculatedPieces []CalculatedPieces
			for r := 0; r < rounds; r++ {
				if manualGeneration == "y" {
					fmt.Print("\nManual Pieces Cost Generation: ")
					fmt.Scanln(&generation)
				} else {
					generation = rand.Intn(12) + 1
				}
				if manualDepth == "y" {
					fmt.Print("\nGame Depth: ")
					fmt.Scanln(&gameDepth)
					pickFirstCost(generation, gameDepth)
				}
				if handmatig == "y" {
					fmt.Print("\nPlayer startposition: ")
					fmt.Scanln(&playerPos)
				} else {
					if randomBoard == "y" {
						playerBoard = rand.Intn(EndBoard)
					}
					playerPos = rand.Intn(53)
				}
				for p := 0; p < amountplayers; p++ {
					if players[p].human {
						var curAnimal string
						var curIdAnimal int
						fmt.Print("Position: ", playerPos, " Player ", players[p].name, " choose animal: ")
						fmt.Scanln(&curAnimal)
						curIdAnimal = getPieceNR(curAnimal)
						calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
						index := 0
						for curIdAnimal != calculatedPieces[index].id {
							index++
						}
						pieces[p] = calculatedPieces[index]
					} else if !players[p].random && !players[p].internal {
						temp := players[p].nodes / gameDepth
						players[p].cd0 = int(temp)
						players[p].cd1 = gameDepth
						moveselect = players[p].moveselect
					} else if !players[p].internal {
						players[p].cd1 = gameDepth
						pieces := int(TotalPieces - 40)
						moveselect = 40 + (rand.Intn(pieces))
					}
					if !players[p].internal {
						calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, players[p].bling, 0, players[p].cd0, players[p].cd1)
					} else {
						calculatedPieces = calculateInternalEngine(playerBoard, playerPos, gameDepth, threads, players[p].inSelection, players[p].inLevel)

						players[p].moveselect = moveselect
					}
					chosenPiece := calculatedPieces[players[p].moveselect].id
					calculatedPieces = calculate(playerBoard, playerPos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
					index := 0
					for chosenPiece != calculatedPieces[index].id {
						index++
					}
					if !players[p].human {
						pieces[p] = calculatedPieces[index]
					}
					scorePosition := getScorePositions(playerBoard, playerPos, gameDepth, pieces[p].id)
					bmNumbers[p] = scorePosition.bm
				}
				fmt.Println("Generation: ", generation, "Board: ", playerBoard+1, " Position: ", playerPos)
				for currentPlayer := range pieces {
					for opponent := 0; opponent < amountplayers; opponent++ {
						if currentPlayer != opponent {
							var result int = whowins012(pieces[currentPlayer].totalCost, bmNumbers[currentPlayer], pieces[opponent].totalCost, bmNumbers[opponent])
							if result == 0 {
								players[currentPlayer].totalPoints += 0.5
							} else if result == 1 {
								players[currentPlayer].totalPoints += 1
							}
						}
					}
					fmt.Println(bmNumbers[currentPlayer], players[currentPlayer].name, pieces[currentPlayer].animal, pieces[currentPlayer].totalCost)
				}
			}
			sort.Slice(players, func(i, j int) bool {
				return ((players[i].totalPoints) > (players[j].totalPoints))
			})
			fmt.Println("\n")
			for o := range players {
				totalpoints := players[o].totalPoints
				fmt.Println(players[o].name, "has a score of", totalpoints, "/", rounds*(amountplayers-1), "\t", (totalpoints/float32(rounds*(amountplayers-1)))*100, "%")
			}
		}

		if choice == 16 {
			var originalBoard int = playerBoard
			var originalPos0 int = playerPos
			var originalPiecesCosts = piecesCosts
			var originalGameDepth uint64 = gameDepth
			var top, board int
			var positionNodes [EndBoard][54]uint64
			var apn uint64 = 0
			var counter uint64 = 0
			fmt.Print("Pieces Costs Generation: ")
			fmt.Scanln(&generation)
			fmt.Scanln(&board)
			fmt.Print("Board (0=all): ")
			fmt.Scanln(&board)
			fmt.Print("Calculate for top: ")
			fmt.Scanln(&top)
			top--
			fmt.Print("Game Depth: ")
			fmt.Scanln(&gameDepth)
			pickFirstCost(generation, gameDepth)
			var b, endB int
			if board == 0 {
				b = 0
				endB = EndBoard
			} else {
				b = board - 1
				endB = board
			}
			for ; b < endB; b++ {
				playerBoard = b
				for pos := 0; pos < 54; pos++ {
					calculatedPieces := calculate(playerBoard, pos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
					topScore := calculatedPieces[top].totalCost
					var j uint64 = gameDepth
					for ; j < gameDepth*uint64(TotalPieces); j += gameDepth {
						cd0, cd1, cd2 = setupEngineByNodes(gameDepth, j)
						calculatedPieces = calculate(playerBoard, pos, threads, gameDepth, false, 0, cd0, cd1)
						if calculatedPieces[0].totalCost == topScore || calculatedPieces[0].totalCost > topScore {
							counter++
							positionNodes[b][pos] = j
							apn += j
							goto nextpos
						}
					}
				nextpos:
				}
			}
			if board == 0 {
				b = 0
				endB = EndBoard
			} else {
				b = board - 1
				endB = board
			}
			for ; b < endB; b++ {
				for j := 0; j < 54; j++ {
					fmt.Println("Board", b+1, " Position", j, " Nodes to get solution in user top:", positionNodes[b][j])
				}
			}
			fmt.Println("Average position nodes: ", (apn / counter))
			playerPos = originalPos0
			gameDepth = originalGameDepth
			playerBoard = originalBoard
			piecesCosts = originalPiecesCosts
			moveselect = 0
		}

		if choice == 17 {
			var originalGameDepth uint64 = gameDepth
			var internal string
			var selection, board, generation int
			var level uint64
			fmt.Print("Pieces Costs Generation: ")
			fmt.Scanln(&generation)
			fmt.Print("Board (0=all): ")
			fmt.Scanln(&board)
			fmt.Print("Game Depth: ")
			fmt.Scanln(&gameDepth)
			pickFirstCost(generation, gameDepth)
			fmt.Print("Internal engine? (y/.): ")
			fmt.Scanln(&internal)
			if internal == "y" {
				selection, level = chooseInternalEngine()
			} else {
				calculateWithInput()
			}
			var b int
			var endB int
			if board == 0 {
				b = 0
				endB = EndBoard
			} else {
				b = board - 1
				endB = board
			}
			for ; b < endB; b++ {
				for pos := 0; pos < 54; pos++ {
					var calculatedPieces []CalculatedPieces
					if internal == "y" {
						calculatedPieces = calculateInternalEngine(b, pos, gameDepth, threads, selection, level)
					} else {
						calculatedPieces = calculate(b, pos, threads, gameDepth, false, 0, cd0, cd1)
					}
					id := calculatedPieces[moveselect].id
					fmt.Print("\nBoard", b+1, " Position", pos, " PICKED: ", calculatedPieces[moveselect].name, " ", calculatedPieces[moveselect].animal, ", bm: ")
					scorePosition := getScorePositions(b, pos, gameDepth, id)
					fmt.Print(scorePosition.bm)
				}
			}
			gameDepth = originalGameDepth
		}

		if choice == 18 {
			var originalBoard int = playerBoard
			var originalPos0 int = playerPos
			var originalGameDepth uint64 = gameDepth
			var originalPiecesCosts = piecesCosts
			var afs uint64 = 0
			var board, generation int
			fmt.Print("Pieces Costs Generation: ")
			fmt.Scanln(&generation)
			fmt.Print("Board (0=all): ")
			fmt.Scanln(&board)
			fmt.Print("Calculate for bm: ")
			fmt.Scanln(&moveselect)
			moveselect--
			fmt.Print("Game Depth: ")
			fmt.Scanln(&gameDepth)
			pickFirstCost(generation, gameDepth)
			var b int
			var endB int
			if board == 0 {
				b = 0
				endB = EndBoard
			} else {
				b = board - 1
				endB = board
			}
			for ; b < endB; b++ {
				playerBoard = b
				for pos := 0; pos < 54; pos++ {
					playerPos = pos
					calculatedPieces := calculate(playerBoard, pos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
					piece := calculatedPieces[moveselect]
					afs += uint64(piece.firstPay)
					fmt.Print("\nBoard ", b+1, " Position ", pos)
					fmt.Print("\n\tBest Animal: ", piece.animal)
					fmt.Print("\n\tTotal Score: ", piece.totalCost)
					fmt.Print("\n\tFirst Score: ", piece.firstPay)
				}
				afs /= 54
				fmt.Print("\nBoard ", b+1, " Average First Score: ", afs)
			}
			playerPos = originalPos0
			gameDepth = originalGameDepth
			playerBoard = originalBoard
			piecesCosts = originalPiecesCosts
			moveselect = 0
		}

		if choice == 19 {
			var originalBoard int = playerBoard
			var originalPos0 int = playerPos
			var originalGameDepth uint64 = gameDepth
			var originalPiecesCosts = piecesCosts
			var board, generation int
			var curAnimal string
			fmt.Print("Pieces Costs Generation: ")
			fmt.Scanln(&generation)
			fmt.Print("Board (0=all): ")
			fmt.Scanln(&board)
			fmt.Print("Calculate for bm: ")
			fmt.Scanln(&moveselect)
			moveselect--
			fmt.Print("Game Depth: ")
			fmt.Scanln(&gameDepth)
			pickFirstCost(generation, gameDepth)
			fmt.Print("Animal: ")
			fmt.Scanln(&curAnimal)
			var idAnimal int = getPieceNR(curAnimal)
			var b int
			var endB int
			if board == 0 {
				b = 0
				endB = EndBoard
			} else {
				b = board - 1
				endB = board
			}
			for ; b < endB; b++ {
				playerBoard = b
				for pos := 0; pos < 54; pos++ {
					playerPos = pos
					calculatedPieces := calculate(playerBoard, pos, threads, gameDepth, false, 0, TotalPieces, gameDepth)
					if calculatedPieces[moveselect].id == idAnimal {
						fmt.Println("Board ", b+1, ", Position ", pos)
					}
				}
			}
			playerPos = originalPos0
			gameDepth = originalGameDepth
			playerBoard = originalBoard
			piecesCosts = originalPiecesCosts
			moveselect = 0
		}

		if choice == 20 {
			var startGeneration, endGeneration, startBoard, endBoard, startPos, endPos int
			var startDepth, endDepth, depthSteps uint64
			fmt.Print("Pieces Costs Start Generation: ")
			fmt.Scanln(&startGeneration)
			fmt.Print("Pieces Costs End Generation: ")
			fmt.Scanln(&endGeneration)
			fmt.Print("Start board: ")
			fmt.Scanln(&startBoard)
			fmt.Print("End board: ")
			fmt.Scanln(&endBoard)
			fmt.Print("Start position: ")
			fmt.Scanln(&startPos)
			fmt.Print("End position: ")
			fmt.Scanln(&endPos)
			fmt.Print("Start gamedepth: ")
			fmt.Scanln(&startDepth)
			fmt.Print("End gamedepth: ")
			fmt.Scanln(&endDepth)
			fmt.Print("In depth steps of: ")
			fmt.Scanln(&depthSteps)
			startBoard--
			endBoard--
			animalCostsHits := make([]float64, 10)
			var totalHits float64
			animalCostsHitsA := make([][]float64, threads)
			totalHitsA := make([]float64, threads)
			for i := range animalCostsHitsA {
				animalCostsHitsA[i] = make([]float64, 10)
			}
			var threadSpread int = (endPos - startPos) / threads
			var newEndPos int = startPos + threadSpread + ((endPos - startPos) % threads)
			startPosA := make([]int, threads)
			endPosA := make([]int, threads)
			startPosA[0] = startPos
			endPosA[0] = (((startPos + newEndPos) % 54) + 54) % 54
			for a := 1; a < threads; a++ {
				startPosA[a] = (((endPosA[a-1] + 1) % 54) + 54) % 54
				endPosA[a] = (((endPosA[a-1] + threadSpread) % 54) + 54) % 54
			}

			var wg sync.WaitGroup
			wg.Add(threads)
			for t := 0; t < threads; t++ {
				go animalCostsMultiThreading(startGeneration, endGeneration, startDepth, endDepth, depthSteps, startBoard, endBoard, startPosA[t], endPosA[t], &animalCostsHitsA[t], &totalHitsA[t], &wg)
			}
			wg.Wait()

			for x := 0; x < threads; x++ {
				totalHits += totalHitsA[x]
				for s := 0; s < 10; s++ {
					animalCostsHits[s] += animalCostsHitsA[x][s]
				}
			}

			for i := range animalCostsHits {
				var percentage float64 = (100 * animalCostsHits[i]) / totalHits
				fmt.Println("Animal Cost ", 5+(i*5), " wint\t", math.Round(percentage*10)/10, "%")
			}
		}

		if choice == 21 {
			var lowestnodes, highestnodes, eDepth uint64
			var trap uint64 = 1
			piecesArray := make([][][]BasicPieces, threads)
			for t := 0; t < threads; t++ {
				piecesArray[t] = make([][]BasicPieces, 20)
				for s := 0; s < 20; s++ {
					piecesArray[t][s] = make([]BasicPieces, TotalPieces)
					for i := 0; i < TotalPieces; i++ {
						piecesArray[t][s][i].id = i
						piecesArray[t][s][i].animal = Animals[i]
						piecesArray[t][s][i].totalCost = 0
					}
				}
			}
			var startGeneration, endGeneration, bStart, bEnd, pStart, pEnd, sTop, eTop int
			var randoms string
			var random uint64 = 0
			/*
				for t := 0; t < threads; t++ {
					for s := 0; s < 20; s++ {
						var i uint64 = 0
						for ; i < TotalPieces; i++ {
							piecesArray[t][s][i].id = i
							piecesArray[t][s][i].animal = Animals[i]
							piecesArray[t][s][i].totalCost = 0
						}
					}
				}
			*/
			fmt.Print("Random position? (y/.): ")
			fmt.Scanln(&randoms)
			if randoms == "y" {
				fmt.Print("How many random positions per depth: ")
				fmt.Scanln(&random)
			} else {
				fmt.Print("Start from board: ")
				fmt.Scanln(&bStart)
				fmt.Print("To board: ")
				fmt.Scanln(&bEnd)
				fmt.Print("Start from position: ")
				fmt.Scanln(&pStart)
				fmt.Print("To position: ")
				fmt.Scanln(&pEnd)
			}
			fmt.Print("Pieces Costs Start Generation: ")
			fmt.Scanln(&startGeneration)
			fmt.Print("Pieces Costs End Generation: ")
			fmt.Scanln(&endGeneration)
			fmt.Print("Start game depth: ")
			fmt.Scanln(&gameDepth)
			fmt.Print("End game depth: ")
			fmt.Scanln(&eDepth)
			fmt.Print("Depth in steps of (1,2,..): ")
			fmt.Scanln(&trap)
			fmt.Print("Top: ")
			fmt.Scanln(&sTop)
			fmt.Print("Bottom: ")
			fmt.Scanln(&eTop)
			fmt.Print("Lowest nodes: ")
			fmt.Scanln(&lowestnodes)
			fmt.Print("Highest nodes: ")
			fmt.Scanln(&highestnodes)
			bStart--
			bEnd--
			sTop--
			eTop--
			currNodes = 0
			calculated := make([]uint64, threads)
			totalNodesArray := make([]uint64, threads)
			var threadSpread int = (pEnd - pStart) / threads
			var newEndPos int = pStart + threadSpread + ((pEnd - pStart) % threads)
			startPos := make([]int, threads)
			endPos := make([]int, threads)
			endPos[0] = (((pStart + newEndPos) % 54) + 54) % 54
			for t := 1; t < threads; t++ {
				startPos[t] = (((endPos[t-1] + 1) % 54) + 54) % 54
				endPos[t] = (((endPos[t-1] + threadSpread) % 54) + 54) % 54
			}

			var wg sync.WaitGroup
			wg.Add(threads)
			var startR uint64 = 0
			var stopR uint64 = (random / uint64(threads)) + (random % uint64(threads))
			for t := 0; t < threads; t++ {
				if randoms != "y" {
					go firstScoreAnimalPerf(startGeneration, endGeneration, gameDepth, eDepth, trap, lowestnodes, highestnodes, bStart, bEnd, startPos[t], endPos[t], sTop, eTop, random, 0, 0, &calculated[t], &piecesArray[t], &totalNodesArray[t], &wg)
				} else {
					go firstScoreAnimalPerf(1, 1, gameDepth, eDepth, trap, lowestnodes, highestnodes, bStart, bEnd, startPos[t], endPos[t], sTop, eTop, random, startR, stopR, &calculated[t], &piecesArray[t], &totalNodesArray[t], &wg)
					startR = stopR
					stopR += (random / uint64(threads))
				}
			}
			wg.Wait()

			var totalNodes uint64
			var totalCalculated uint64
			pieces := make([][]BasicPieces, 20)
			for s := 0; s < 20; s++ {
				pieces[s] = make([]BasicPieces, TotalPieces)
				for p := 0; p < TotalPieces; p++ {
					pieces[s][p].id = p
					pieces[s][p].animal = Animals[p]
					pieces[s][p].totalCost = 0
				}
			}
			for x := 0; x < threads; x++ {
				totalNodes += totalNodesArray[x]
				totalCalculated += calculated[x]
				for s := 0; s < 20; s++ {
					for p := 0; p < int(TotalPieces); p++ {
						id := piecesArray[x][s][p].id
						pieces[s][id].totalCost += piecesArray[x][s][p].totalCost
					}
				}
			}
			for s := 0; s < 20; s++ {
				sort.Slice(pieces[s], func(i, j int) bool {
					return (pieces[s][i].totalCost) > (pieces[s][j].totalCost)
				})
			}

			fmt.Println("\nTotal Nodes: ", totalNodes)
			fmt.Println("Total position calculations: ", totalCalculated)
			for s := 0; s < 20; s++ {
				fmt.Println("\nField with start score: ", s+1)
				var current int = 1
				for i := range pieces[s] {
					if pieces[s][i].totalCost > 0 {
						fmt.Println(current, pieces[s][i].animal, pieces[s][i].totalCost)
						current++
					}
				}
			}
		}

		if choice == 22 {
			getAnimals()
		}

		if choice == 23 {
			var generations int
			fmt.Print("Amount of generations: ")
			fmt.Scanln(&generations)
			createCostsGenerations(generations)
		}
	}
}
