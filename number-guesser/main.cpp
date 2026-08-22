#include <chrono>
#include <iostream>
#include <limits>
#include <random>
#include <thread>

using namespace std;

constexpr int MIN = 0;
constexpr int MAX = 1000;

int get_random_number() {
  random_device rd;
  mt19937 gen(rd());
  uniform_int_distribution<int> dist(MIN, MAX);
  return dist(gen);
}

int main() {
  cout << "Welcome to the number guessing game! \n";
  cout << "First computer can guess the number, please wait! \n";

  // sleep for 1 second for like computer guessing
  this_thread::sleep_for(chrono::seconds(1));

  // random number generation
  int rn = get_random_number();

  cout << "Done \n";
  int user_count = 0;

  while (true) {
    int guess;
    cout << "your guess: ";
    cin >> guess;
    user_count++;

    if (guess == rn) {
      cout << "congratulations!, your total guessess is " << user_count << "\n";
      break;
    } else if (guess > rn) {
      cout << "lower \n";
    } else {
      cout << "higher \n";
    }
  }

  cout << "Now computer turn, choose the number between 1 - 1000 \n";
  cout << "Once you have choose and press enter key to continue... \n";

  cin.ignore(numeric_limits<streamsize>::max(), '\n');
  cin.get();

  int computer_count = 0;
  int secret_number = get_random_number();
  int guess = secret_number;
  int lower = MIN;
  int higher = MAX;
  bool first = true;
  int wrong_key_press_count = 0;

  while (true) {
    cout << "my guess: " << guess << '\n';
    cout << "is that higher or lower or correct (press h or l or c)";
    char ui;
    cin >> ui;

    if (!first) {
      if (guess == secret_number) {
        cout << "sorry f*er. you cheat. otherwise you gave h or l wrongly, "
                "Thanks"
             << "\n";
        goto end_of_loop;
      }
    }
    computer_count++;

    switch (ui) {
    case 'h':
      lower = guess;
      guess = (higher - guess) / 2 + guess;
      break;
    case 'l':
      higher = guess;
      guess = (guess - lower) / 2 + lower;
      break;
    case 'c':
      cout << "Thanks" << '\n';
      cout << "I got it in " << computer_count << " attempts \n";
      goto end_of_loop;
    default: {
      cout << "wrong key press again..";
      wrong_key_press_count++;
      first = true;
    }
    }
    first = false;

    if (wrong_key_press_count >= 3) {
      goto end_of_loop;
    }
  }

end_of_loop:

  if (wrong_key_press_count >= 3) {
    cout << "sorry, maximum retry attempts was completed" << '\n';
    return 0;
  }

  if (user_count < computer_count) {
    cout << "User wins, congratulations" << "\n";
  }

  if (user_count > computer_count) {
    cout << "Computer wins, try again" << "\n";
  }

  return 0;
}
