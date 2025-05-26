# Go Web Crawler

This program is a simple web crawler built with Go. It crawls web pages starting from a given URL, staying within the same domain. It keeps track of the number of times each page is visited and can limit the number of concurrent requests and the total number of pages crawled.

## Installation

1.  **Install Go:** Ensure you have Go (version 1.x or later) installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2.  **Clone the repository (Optional):** If you haven't already, clone this repository to your local machine.
3.  **Build the executable:** Navigate to the project's root directory in your terminal and run the following command to build the crawler:
    ```bash
    go build
    ```
    This will create an executable file named `crawler` (or `crawler.exe` on Windows) in the project directory.

## Usage

Once built, you can run the crawler from your terminal using the following command structure:

```bash
./crawler <URL> <maxConcurrency> <maxPages>
```

**Arguments:**

*   `<URL>`: The starting URL for the crawl (e.g., `https://example.com`). The crawler will only visit pages within the same domain as this initial URL.
*   `<maxConcurrency>`: An integer specifying the maximum number of concurrent HTTP requests the crawler can make. This helps to control the load on the server being crawled.
*   `<maxPages>`: An integer specifying the maximum number of unique pages the crawler will visit.

**Example:**

To crawl `https://golang.org` with a maximum of 10 concurrent requests and up to 500 pages:

```bash
./crawler https://golang.org 10 500
```

**Output:**

The crawler will print the URLs it crawls in real-time. Once the crawl is complete or the `maxPages` limit is reached, it will display:

1.  A list of all unique URLs found and the number of times each was visited.
2.  The total time taken for the crawl.

Example output snippet:

```
Crawling: https://golang.org/
Got HTML
Crawling: https://golang.org/doc/
Got HTML
...
Results:
https://golang.org/: 1
https://golang.org/doc/: 1
...
Crawl completed - Elapsed time:  XXmXX.XXXs
```
