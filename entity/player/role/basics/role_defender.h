/***
 * Maracatronics Robotics
 * Federal University of Pernambuco (UFPE) at Recife
 * http://www.maracatronics.com/
 *
 * This file is part of Armorial project.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 ***/

#ifndef ROLE_DEFENDER_H
#define ROLE_DEFENDER_H

#include <entity/player/behaviour/mrcbehaviours.h>
#include <entity/player/role/role.h>

class Role_Defender : public Role
{
private:
    // Behaviours
    Behaviour_Barrier *_bh_bar;

    // Behaviours Enum
    enum{
        BEHAVIOUR_BARRIER
    };

    // Inherited functions
    void configure();
    void run();

    // Parameters
    int _barrierId;
    bool _barrierCanTakeout;
    char _barrierSide;

    // Auxiliary Functions
    float distanceFromMidGoalShoot(char side);

public:
    Role_Defender();
    void initializeBehaviours();
    QString name();

    void setBarrierId(int barrierId) { _barrierId = barrierId; }
    void setBarrierCanTakeout(bool barrierCanTakeout) { _barrierCanTakeout = barrierCanTakeout; }
    void setBarrierSide(char barrierSide) { _barrierSide = barrierSide; }
};

#endif // ROLE_DEFENDER_H
