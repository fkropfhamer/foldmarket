# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     Foldmarket.Repo.insert!(%Foldmarket.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.

alias Foldmarket.Repo
alias Foldmarket.Accounts.User

if Mix.env() == :dev do
  Repo.insert!(%User{
    email: "user@example.com",
    hashed_password: User.hash_password("1234"),
  })
end
