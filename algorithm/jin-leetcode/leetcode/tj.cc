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

const string outconfigpath="c.txt";
vector<string> v;

vector<string> subStrToVec(string str, char sep)
{
    vector<string> vecArr;
    int flagSub = 0;
    for (int i = 0; i < str.length(); i++)
    {
        if (str[i] == sep)
        {
            string temp = str.substr(flagSub, i - flagSub);
            vecArr.push_back(temp);
            v.push_back(temp);
            flagSub = i + 1;
        }
    }
    string temp = str.substr(flagSub, str.length());
    vecArr.push_back(temp);
    return vecArr;
}
void LoadConfigLine(string str)
{
    vector<string> vecArr = subStrToVec(str, ',');
   
}
bool LoadLocalConfig()
{
    
    std::ifstream ifs(outconfigpath, std::ios::in);
    if (!ifs.is_open())
    {
        cout << endl;
        cout << outconfigpath << "Failed to open configpath.\n";
        ofstream outftxt;
        outftxt.open(outconfigpath, ios::out | ios::trunc);
        outftxt.close();
    }
    else
    {
        std::string buf;
        while (getline(ifs, buf))
        {
            cout << buf << "   " << buf.length() << endl;
            if (buf.size() > 0)
            {
                LoadConfigLine(buf);
            }
        }
        ifs.close();
    }
    return true;
}

std::string& trim(std::string &s) 
{
    if (s.empty()) 
    {
        return s;
    }
 
    s.erase(0,s.find_first_not_of(" "));
    s.erase(s.find_last_not_of(" ") + 1);
    return s;
}

int main() {
   LoadLocalConfig();
   int i=0;
   cout<<"----------"<<v.size()<<endl;
    for(auto s:v){
        cout<<++i<<"--"<<trim(s)<<endl;
    }
    return 0;
}




