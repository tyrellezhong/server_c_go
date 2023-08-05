#pragma once
#include <string>
#include "../includes/Gun.h"

class Solider {
public: 
    Solider(std::string name);
    ~Solider();
    void AddBulletToGun(int num);
    void AddGun(Gun *ptr_gun);
    bool fire();
private:
    std::string _name;
    Gun *_ptr_gun;
};