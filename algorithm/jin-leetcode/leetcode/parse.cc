#include <iostream>
#include <fstream>
#include <string>
#include <set>
#include <iostream>
#include <iomanip>
#include <vector>
#include <string.h>
#include <stdio.h>
#include <cstddef>         // std::size_t
using namespace std;


int main() {
    ifstream file("a.txt");
    set<int> st;
    string line;
    string s="      ";
        ofstream out_txt_file;
    out_txt_file.open("b.txt");
    while (getline(file, line)) {
       if(line.find("#")== std::string::npos){
            s += line+"     ";
       }else{
        string v = line+ s;
        s="     ";
        if(v.size()>0){
            // char* s = new char[v.size() + 1];
            // strcpy(s, v.c_str());
            // char* p = strtok(s, " ");
            // vector<string> words;
            // while(p) {
            //     words.push_back(p);
            //     p = strtok(NULL, " ");
            // }
                istringstream ss(v);
                vector<string> words;
                string word;
                while(ss >> word) {
                    words.push_back(word);
                }
            out_txt_file<< setw(20) << setfill(' ');
            for(string x : words) {
                out_txt_file <<x;
            }
            out_txt_file <<endl;
            
        }

       }
        
    }
out_txt_file.close();
;
    return 0;
}




