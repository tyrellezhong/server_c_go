#include "../includes/Soldier.h"

Solider:: Solider(std::string name) {
    this->_name = name;
    this->_ptr_gun = nullptr;
}
void Solider::AddGun(Gun *ptr_gun) {
    this->_ptr_gun = ptr_gun;
}
void Solider::AddBulletToGun(int num) {
    this->_ptr_gun->AddBullet(num);
}
bool Solider::fire() {
    return this->_ptr_gun->Shoot();
}
Solider::~Solider() {
    if (this->_ptr_gun == nullptr) {
        return;
    }
    delete this->_ptr_gun;
    this->_ptr_gun = nullptr;
}