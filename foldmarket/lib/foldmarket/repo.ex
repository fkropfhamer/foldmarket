defmodule Foldmarket.Repo do
  use Ecto.Repo,
    otp_app: :foldmarket,
    adapter: Ecto.Adapters.Postgres
end
