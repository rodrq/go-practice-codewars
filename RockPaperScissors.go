/* Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"
*/

package kata

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

func Rps(a, b string) string {
	if a == b {
		return "Draw!"
	}
	if m[a] == b {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}

func RpsWut(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	a := p1 + p2
	switch a {
	case "paperrock", "scissorspaper", "rockscissors":
		return "Player 1 won!"
	default:
		return "Player 2 won!"
	}
}

func RpsXD(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}
	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}
	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}
	return "Player 2 won!"
}
