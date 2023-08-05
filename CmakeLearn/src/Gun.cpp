#include "../includes/Gun.h"
#include <iostream>
void Gun::AddBullet(int bullet_num){
    this->_bullet_count += bullet_num;
}
bool Gun::Shoot() {
    if (this->_bullet_count <= 0) {
        std::cout << "There is no bullet !" << std::endl;
        return false;
    }
    this->_bullet_count -= 1;
    std::cout << "shoot sueecssfully !" << std::endl;
    return true;
}