defmodule Main do
  def sort(str) do
    str
    |> String.to_charlist()
    |> Enum.sort(&(&1 < &2))
  end

  def valid_anagram(s1, s2) do
    sort(s1) == sort(s2)
  end
end

IO.inspect(Main.valid_anagram("anagram", "nagaram"))
IO.inspect(Main.valid_anagram("anagr", "nagaram"))
