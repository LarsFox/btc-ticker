# BTC Price ticker

The tickers supply the price of BTC every two seconds, while the streamer reads the tickers’ prices and stores them one by one to print the calculated *fair price* every minute.

Once the price is printed, its source data is deleted, while all of the delayed messages are ignored, since time travelling should be implemented in a separate package. This also prevents storage from overflowing.

Use `make` to launch the live demo. Make sure to speed it up by setting the streamer duration to, say, 5 seconds.

I do not have a degree in BTC *fair prices* algorhythms, however, I implemented the two most commonly used functions to calculate something average: mean and median. Feel free to try them out.

## To improve

1. Move `PriceStreamSubscriber` to the package `streamer`, where it is used.
2. Move index key calculating outside of the `storage` package.
