# Game Programming

A collection of small games built for fun and learning.

## 🎮 Game 1: Number Guesser

A classic number guessing game with a twist — you play **against the computer**.

### How to Play

1. **Compile the game** (requires a C++ compiler):
   ```bash
   cd number-guesser
   g++ main.cpp -o a.out
   ```

2. **Run it**:
   ```bash
   ./a.out
   ```

3. **Round 1 — You guess**: The computer picks a secret number between 0 and 1000. Enter your guesses, and it will tell you `higher` or `lower` until you get it right.

4. **Round 2 — Computer guesses**: Now you pick a number between 0 and 1000. The computer starts guessing — you respond with:
   - `h` if your number is higher
   - `l` if your number is lower
   - `c` if the computer got it correct

5. Whoever used fewer guesses wins. Try to beat its binary search... if you can. 😉

> Built on top of a powerful binary search — see [number-guesser/readme.md](number-guesser/readme.md) for more.

---

## 🔥 Game 2: Super Card *(currently cooking 👨‍🍳)*

Something bigger is brewing in the kitchen...

**Super Card** is a multiplayer card game built in Go with real-time gameplay over WebSockets. Think random superstars, live rooms, and head-to-head action in the browser.

It's still on the stove — features are being added, flavors are being tuned. Stay tuned for the full recipe and play instructions when it's ready to serve. 🍳
