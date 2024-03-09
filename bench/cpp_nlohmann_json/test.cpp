#include <iostream>
#include <fstream>
#include <string>
#include "json.hpp"

using namespace std;

using json = nlohmann::json;

json read_json_file(const std::string& filename)
{
    std::ifstream input_file(filename);
    if (!input_file.is_open())
    {
        std::cerr << "Failed to open file '" << filename << "'" << std::endl;
        return json();
    }

    // Read the contents of the file
    std::string file_contents((std::istreambuf_iterator<char>(input_file)),
                              std::istreambuf_iterator<char>());

    // Parse the JSON data and return it
    return json::parse(file_contents);
}

string read_file(const std::string& filename)
{
    std::ifstream input_file(filename);
    if (!input_file.is_open())
    {
        std::cerr << "Failed to open file '" << filename << "'" << std::endl;
        return json();
    }

    // Read the contents of the file
    std::string file_contents((std::istreambuf_iterator<char>(input_file)),
                              std::istreambuf_iterator<char>());

    // Parse the JSON data and return it
    return std::move(file_contents);
}

int main(int argc, char* argv[])
{
    if (argc != 2)
    {
        std::cout << "Usage: " << argv[0] << " input_file" << std::endl;
        return 1;
    }

    const std::string file = read_file(argv[1]);

    // Repeat the process for 100 times
    for (int i = 0; i < 800; ++i)
    {
        std::cout << "=============== Iteration " << i + 1 << " ===============" << std::endl;
        // Read and parse the JSON data from the file
        json data = json::parse(file);

        // Print the parsed data to stdout
        std::cout << data.dump(4)[100] << std::endl;
    }

    return 0;
}
