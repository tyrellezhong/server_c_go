#pragma once
#include <string>

class Gun {
public:
    Gun(std::string type) {
        this->_bullet_count = 0;
        this->_type = type;
    }
    void AddBullet(int bullet_num);
    bool Shoot();
private:
    int _bullet_count;
    std::string _type;
};