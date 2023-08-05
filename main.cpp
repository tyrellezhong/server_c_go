#include "includes/Gun.h"
#include "includes/Soldier.h"

void SoliderShoot() {
    Solider sanduo("xusanduo");
    sanduo.AddGun(new Gun("AK47"));
    sanduo.AddBulletToGun(20);
    sanduo.fire();
}
int main() {
    SoliderShoot();
    return 0;
}