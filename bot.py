import TwitterAPI

print('Logging in...')
api = TwitterAPI.TwitterAPI('nFAyRc7404OlFnGrzNyglJq1d', 'RYnMhEG49dH0O2Pgz0Jw4IIqViBL2rZ0DbkLnSTpETjZT12ENN', '349591688-uYVZ4ricRGjjoa4VAHB29S52MGG0J9Orat6gA7Lx', 'pKtZrpPjMojJYBr8OhN1edZA7UwE8ONDAZZIcHM56iA5h')
print('Done.')
r = api.request('statuses/filter', {'follow':'898994539'})
for item in r:
    print(item['text'] if 'text' in item else item)